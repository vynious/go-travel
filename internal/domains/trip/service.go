package trip

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type Service struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewTripService(repository *repo.Repository) *Service {
	return &Service{
		repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *Service) CreateNewTrip(ctx context.Context, title string, country string, startDate time.Time, endDate time.Time) (db.Trip, error) {

	params := db.CreateTripParams{
		Title:     title,
		Country:   country,
		StartDate: startDate,
		EndDate:   endDate,
	}

	trip, err := s.repository.Queries.CreateTrip(ctx, params)

	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to create trip: %w", err)
	}
	return trip, nil
}

func (s *Service) GetAllTrips(ctx context.Context) ([]db.Trip, error) {
	trips, err := s.repository.Queries.ListTrips(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get trips: %w", err)
	}
	return trips, nil
}

func (s *Service) GetTripById(ctx context.Context, id int64) (db.Trip, error) {
	trip, err := s.repository.Queries.GetTrip(ctx, id)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to get trip: %w", err)
	}
	return trip, nil
}

func (s *Service) UpdateTripTitle(ctx context.Context, id int64, title string) (db.Trip, error) {
	updateParams := db.UpdateTripTitleParams{
		ID:    id,
		Title: title,
	}
	trip, err := s.repository.Queries.UpdateTripTitle(ctx, updateParams)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to update title: %w", err)
	}
	return trip, nil
}

func (s *Service) UpdateTripCountry(ctx context.Context, id int64, country string) (db.Trip, error) {
	updateParams := db.UpdateTripCountryParams{
		ID:      id,
		Country: country,
	}
	trip, err := s.repository.Queries.UpdateTripCountry(ctx, updateParams)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to update country: %w", err)
	}
	return trip, nil
}

func (s *Service) UpdateTripStartDate(ctx context.Context, id int64, startDate time.Time) (db.Trip, error) {
	updateParams := db.UpdateTripStartDateParams{
		ID:        id,
		StartDate: startDate,
	}
	trip, err := s.repository.Queries.UpdateTripStartDate(ctx, updateParams)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to update start date: %w", err)
	}
	return trip, nil
}

func (s *Service) UpdateTripEndDate(ctx context.Context, id int64, endDate time.Time) (db.Trip, error) {
	updateParams := db.UpdateTripEndDateParams{
		ID:      id,
		EndDate: endDate,
	}
	trip, err := s.repository.Queries.UpdateTripEndDate(ctx, updateParams)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to update end date: %w", err)
	}
	return trip, nil
}

func (s *Service) DeleteTripById(ctx context.Context, id int64) (db.Trip, error) {
	user, err := s.repository.Queries.DeleteTrip(ctx, id)
	if err != nil {
		return db.Trip{}, fmt.Errorf("unable to delete trip: %w", err)
	}
	return user, nil
}
