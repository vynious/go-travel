package connections

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type ConnectionService struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewConnectionService(repo *repo.Repository) *ConnectionService {
	return &ConnectionService{
		repository: repo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *ConnectionService) CreateConnection(ctx context.Context, a string, b string) (db.Connection, error) {
	params := db.CreateConnectionParams{
		PartyA: a,
		PartyB: b,
	}
	conn, err := s.repository.Queries.CreateConnection(ctx, params)
	if err != nil {
		return db.Connection{}, fmt.Errorf("[s] failed to create connection: %w", err)
	}
	return conn, nil
}

func (s *ConnectionService) GetUserConnections(ctx context.Context, userId string) ([]db.GetConnectionsByUserIdRow, error) {
	conns, err := s.repository.Queries.GetConnectionsByUserId(ctx, userId)
	if err != nil {
		return []db.GetConnectionsByUserIdRow{}, fmt.Errorf("[s] failed to retreive user connections")
	}
	return conns, nil
}

func (s *ConnectionService) DeleteConnection(ctx context.Context, a string, b string) (db.Connection, error) {
	params := db.DeleteConnectionByUserIdParams{
		PartyA: a,
		PartyB: b,
	}
	conn, err := s.repository.Queries.DeleteConnectionByUserId(ctx, params)
	if err != nil {
		return db.Connection{}, fmt.Errorf("[s] failed to delete connection: %w", err)
	}
	return conn, nil
}
