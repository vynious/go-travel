package user

import (
	"context"
	"fmt"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type Service struct {
	repository *Repository
	timeout    time.Duration
}

func NewUserService(repository *Repository) *Service {
	return &Service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *Service) CreateNewUser(ctx context.Context, name string, username string, email string) (db.User, error) {
	registrationArg := db.CreateUserParams{
		Name:     name,
		Username: username,
		Email:    email,
	}
	user, err := s.repository.queries.CreateUser(ctx, registrationArg)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to create user: %w", err)
	}
	return user, nil
}

func (s *Service) GetUserById(ctx context.Context, id int64) (db.User, error) {
	user, err := s.repository.queries.GetUser(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := s.repository.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	user, err := s.repository.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

// VerifyUserExistence takes in
func (s *Service) VerifyUserExistence(ctx context.Context, email string, password string) (bool, error) {
	// call the firebase auth service
	// if true

	user, err := s.repository.queries.GetUserByEmail(ctx, email)

	if err != nil {
		return false, err
	}
	return true, nil
}
