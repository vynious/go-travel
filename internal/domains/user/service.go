package user

import (
	"context"
	"database/sql"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type UserService struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewUserService(repository *repo.Repository) *UserService {
	return &UserService{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *UserService) CreateNewUser(ctx context.Context, uid string, name string, username string, email string) (db.User, error) {
	registrationArg := db.CreateUserParams{
		ID:       uid,
		Name:     name,
		Username: username,
		Email:    email,
	}
	user, err := s.repository.Queries.CreateUser(ctx, registrationArg)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to create user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetAllUser(ctx context.Context) ([]db.User, error) {
	users, err := s.repository.Queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get all user: %w", err)
	}
	return users, nil
}

func (s *UserService) GetUserById(ctx context.Context, id string) (db.User, error) {
	user, err := s.repository.Queries.GetUser(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := s.repository.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (db.User, error) {
	user, err := s.repository.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUserPictureById(ctx context.Context, id string, url string) (db.User, error) {
	profilePicture := sql.NullString{String: url, Valid: url != ""}

	updateParams := db.UpdateUserProfilePictureParams{
		ID:             id,
		ProfilePicture: profilePicture,
	}
	user, err := s.repository.Queries.UpdateUserProfilePicture(ctx, updateParams)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to update profile picture: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUserEmailById(ctx context.Context, id string, email string) (db.User, error) {
	updateParams := db.UpdateUserEmailParams{
		ID:    id,
		Email: email,
	}
	user, err := s.repository.Queries.UpdateUserEmail(ctx, updateParams)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to update email: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUserUsernameById(ctx context.Context, id string, username string) (db.User, error) {
	updateParams := db.UpdateUserUsernameParams{
		ID:       id,
		Username: username,
	}
	user, err := s.repository.Queries.UpdateUserUsername(ctx, updateParams)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to update username: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUserNameById(ctx context.Context, id string, name string) (db.User, error) {
	updateParams := db.UpdateUserNameParams{
		ID:   id,
		Name: name,
	}
	user, err := s.repository.Queries.UpdateUserName(ctx, updateParams)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to update name: %w", err)
	}
	return user, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, id string) (db.User, error) {
	user, err := s.repository.Queries.DeleteUser(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to delete user: %w", err)
	}
	return user, nil
}

// VerifyUserExistence takes in
func (s *UserService) VerifyUserExistence(ctx context.Context, email string, password string) (bool, error) {
	// call the firebase auth service
	// if true
	_, err := s.repository.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
