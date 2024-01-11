// interacting with the database [data access layer]

package db

import (
	"database/sql"
	sqlc "github.com/vynious/go-travel/internal/db/sqlc"
)

type Repository struct {
	Queries *sqlc.Queries
	DB      *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB:      db,
		Queries: sqlc.New(db),
	}
}
