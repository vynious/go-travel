package auth

import (
	"github.com/joho/godotenv"
	"log"
)

func NewAuth() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load .env: %v", err)
	}

}
