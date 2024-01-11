package trip

import (
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type StartTripRequest struct {
	Title     string
	Country   string
	StartDate time.Time
	EndDate   time.Time
}

type UpdateTripDetailRequest struct {
	Title     *string
	Country   *string
	StartDate *time.Time
	EndDate   *time.Time
}

type TripDetailResponse struct {
	Trip db.Trip
}

type AllTripDetailResponse struct {
	Trips []db.Trip
}
