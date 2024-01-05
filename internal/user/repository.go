// interacting with the database [data access layer]

package user

import (
	"database/sql"
	sqlc "github.com/vynious/go-travel/internal/db/sqlc"
)

type Repository struct {
	queries *sqlc.Queries
	db      *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db:      db,
		queries: sqlc.New(db),
	}
}
