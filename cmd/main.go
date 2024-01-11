package main

import (
	"database/sql"
	"github.com/vynious/go-travel/internal/domains/auth"
	"net/http"
)

type App struct {
	router         http.Handler
	rdb            *sql.DB
	firebaseClient *auth.Client
}

func main() {

}
