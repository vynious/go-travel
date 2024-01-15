package travel_entry

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type TravelEntryService struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewTravelEntryService(repo *repo.Repository) *TravelEntryService {
	return &TravelEntryService{
		repository: repo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *TravelEntryService) CreateNewTravelEntry(ctx context.Context, uid string, tid int64, location string, description string) (db.TravelEntry, error) {

	params := db.CreateTravelEntryParams{
		UserID:      uid,
		TripID:      tid,
		Location:    location,
		Description: description,
	}

	entry, err := s.repository.Queries.CreateTravelEntry(ctx, params)
	if err != nil {
		return db.TravelEntry{}, fmt.Errorf("[service] failed to create travel entry: %w", err)
	}
	return entry, nil
}

func (s *TravelEntryService) GetTravelEntryById(ctx, eid int64) (db.TravelEntry, error) {

}

func (s *TravelEntryService) GetTravelEntriesByTripId(ctx, tid int64) ([]db.TravelEntry, error) {

}

func (s *TravelEntryService) GetTravelEntriesByUserId(ctx, uid string) ([]db.TravelEntry, error) {

}

func (s *TravelEntryService) UpdateTravelEntryLocationById(ctx, eid int64, location string) (db.TravelEntry, error) {

}
func (s *TravelEntryService) UpdateTravelEntryDescriptionById(ctx, eid int64, description string) (db.TravelEntry, error) {

}

func (s *TravelEntryService) DeleteTravelEntryById(ctx, eid int64) (db.TravelEntry, error) {

}
