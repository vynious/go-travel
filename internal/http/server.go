package http

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/vynious/go-travel/internal/db"
	auth "github.com/vynious/go-travel/internal/domains/auth"
	"github.com/vynious/go-travel/internal/domains/trip"
	"github.com/vynious/go-travel/internal/domains/user"
	"github.com/vynious/go-travel/pkg"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	router         http.Handler
	config         Config
	rdb            *sql.DB
	firebaseClient *auth.Client
}

func NewApp() (*App, error) {
	dsn := "postgresql://shawntyw:shawntyw@localhost/godb?sslmode=disable"

	// Open the database
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Open Firebase Auth
	fireClient, err := auth.NewFirebaseClient()
	if err != nil {
		return nil, fmt.Errorf("error starting firebase: %w", err)
	}

	repo := db.NewRepository(database)

	userService := user.NewUserService(repo)
	userHandler := user.NewUserHandler(userService, fireClient)

	tripService := trip.NewTripService(repo)
	tripHandler := trip.NewTripHandler(tripService)

	app := &App{
		router:         InitRouter(userHandler, tripHandler),
		rdb:            database,
		config:         LoadConfig(),
		firebaseClient: fireClient,
	}

	return app, nil
}

func (a *App) Start() error {
	// Create a channel to listen for an interrupt or termination signal from the OS.
	// This is used for graceful server shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Create a server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort), // Assuming cfg.ServerPort is the port number
		Handler: a.router,                                // Your HTTP handlers
	}

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			pkg.Log.Error("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	<-quit
	pkg.Log.Info("Shutting down server.")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	}

	// Call App.Close() to clean up other resources like Redis
	if err := a.Close(); err != nil {
		log.Printf("Error during application cleanup: %v", err)
	}

	pkg.Log.Info("Server and resources closed successfully.")
	return nil
}

func (a *App) Close() error {
	if err := a.rdb.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}
	return nil
}
