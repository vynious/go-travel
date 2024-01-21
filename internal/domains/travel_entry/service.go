package travel_entry

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"strconv"
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

func (s *TravelEntryService) GetTravelEntryById(ctx context.Context, eid int64) (db.TravelEntry, error) {
	entry, err := s.repository.Queries.GetTravelEntryById(ctx, eid)
	if err != nil {
		return db.TravelEntry{}, fmt.Errorf("[service] failed to get travel entry: %w", err)
	}
	return entry, nil
}

func (s *TravelEntryService) GetTravelEntriesByTripId(ctx context.Context, tid int64) ([]db.TravelEntry, error) {
	entries, err := s.repository.Queries.GetAllTravelEntryByTripId(ctx, tid)
	if err != nil {
		return []db.TravelEntry{}, fmt.Errorf("[service] failed to get travel entries: %w", err)
	}
	return entries, nil
}

func (s *TravelEntryService) GetTravelEntriesByUserIdAndTripId(ctx context.Context, uid string, tid int64) ([]db.TravelEntry, error) {
	params := db.GetAllTravelEntryByUserIdAndTripIdParams{
		TripID: tid,
		UserID: uid,
	}
	entries, err := s.repository.Queries.GetAllTravelEntryByUserIdAndTripId(ctx, params)
	if err != nil {
		return []db.TravelEntry{}, fmt.Errorf("[service] failed to get travel entries: %w", err)
	}
	return entries, nil
}

func (s *TravelEntryService) UpdateTravelEntryLocationById(ctx context.Context, eid int64, location string) (db.TravelEntry, error) {
	params := db.UpdateTravelEntryLocationParams{
		Location: location,
		ID:       eid,
	}
	entry, err := s.repository.Queries.UpdateTravelEntryLocation(ctx, params)
	if err != nil {
		return db.TravelEntry{}, fmt.Errorf("[service] failed to update travel entry: %w", err)
	}
	return entry, nil
}
func (s *TravelEntryService) UpdateTravelEntryDescriptionById(ctx context.Context, eid int64, description string) (db.TravelEntry, error) {
	params := db.UpdateTravelEntryDescriptionParams{
		Description: description,
		ID:          eid,
	}
	entry, err := s.repository.Queries.UpdateTravelEntryDescription(ctx, params)
	if err != nil {
		return db.TravelEntry{}, fmt.Errorf("[service] failed to update travel entry: %w", err)
	}
	return entry, nil
}

func (s *TravelEntryService) DeleteTravelEntryById(ctx context.Context, eid int64) (db.TravelEntry, error) {
	entry, err := s.repository.Queries.DeleteTravelEntry(ctx, eid)
	if err != nil {
		return db.TravelEntry{}, fmt.Errorf("[service] failed to delete travel entry: %w", err)
	}
	return entry, nil

}

func generateS3Key(eid int64, filename string) string {
	strEID := strconv.FormatInt(eid, 10)
	key := strEID + "_" + filename
	return key
}
