package user_trip

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type UserTripService struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewUserTripService(repository *repo.Repository) *UserTripService {
	return &UserTripService{
		repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *UserTripService) CreateNewUserTrip(ctx context.Context, uid string, tid int64) (db.UserTrip, error) {

	params := db.CreateUserTripParams{
		TripID: tid,
		UserID: uid,
	}

	ut, err := s.repository.Queries.CreateUserTrip(ctx, params)
	if err != nil {
		return db.UserTrip{}, fmt.Errorf("unable to assign user to trip: %w", err)
	}
	return ut, nil
}

// GetUserTripByTripId : Getting the users on a specific trip
func (s *UserTripService) GetUserTripByTripId(ctx context.Context, tid int64) ([]db.UserTrip, error) {

	uts, err := s.repository.Queries.GetUserTripsByTripId(ctx, tid)
	if err != nil {
		return []db.UserTrip{}, fmt.Errorf("unable to get users on trip: %w", err)
	}

	return uts, nil
}

// GetUserTripByUserId : Getting the trips of a specific user
func (s *UserTripService) GetUserTripByUserId(ctx context.Context, uid string) ([]db.UserTrip, error) {

	uts, err := s.repository.Queries.GetUserTripsByUserId(ctx, uid)
	if err != nil {
		return []db.UserTrip{}, fmt.Errorf("unable to get trips of user: %w", err)
	}

	return uts, nil
}

func (s *UserTripService) DeleteUserTripById(ctx context.Context, uid string, tid int64) (db.UserTrip, error) {

	params := db.DeleteUserTripParams{
		UserID: uid,
		TripID: tid,
	}

	ut, err := s.repository.Queries.DeleteUserTrip(ctx, params)
	if err != nil {
		return db.UserTrip{}, fmt.Errorf("unable to delete user_trip assignment: %w", err)
	}

	return ut, err
}
