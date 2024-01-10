package main

import (
	"database/sql"
	"github.com/vynious/go-travel/internal/auth"
	"net/http"
)

type App struct {
	router         http.Handler
	rdb            *sql.DB
	firebaseClient *auth.Client
}

func main() {

}
