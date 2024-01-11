package main

import (
	"github.com/vynious/go-travel/internal/http"
	"log"
)

func main() {
	app, err := http.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application
	if err := app.Start(); err != nil {
		log.Fatalf("failed to start application: %v", err)
	}
}
