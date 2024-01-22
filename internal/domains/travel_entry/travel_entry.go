package travel_entry

import (
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	db "github.com/vynious/go-travel/internal/db/sqlc"
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
	SignedUrls  []*v4.PresignedHTTPRequest
}
