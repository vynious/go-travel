package travel_entry

import (
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"github.com/vynious/go-travel/internal/domains/media"
)

type NewTravelEntryRequest struct {
	UserId      string
	TripId      int64
	Location    string
	Description string
}

type TravelEntryDetailResponse struct {
	TravelEntry db.TravelEntry
}

type TravelEntriesDetailResponse struct {
	TravelEntries []db.TravelEntry
}

type UpdateTravelEntryRequest struct {
	Location    *string
	Description *string
}

type TravelEntryDetailWithMediaResponse struct {
	TravelEntry db.TravelEntry
	SignedMedia []*media.MediaResponse
}
