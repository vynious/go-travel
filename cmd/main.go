package main

import (
	"github.com/joho/godotenv"
	"github.com/vynious/go-travel/internal/http"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app, err := http.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application
	if err := app.Start(); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}
