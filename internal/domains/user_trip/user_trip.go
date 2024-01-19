package user_trip

import db "github.com/vynious/go-travel/internal/db/sqlc"

type AddUserTripRequest struct {
	UserId string
}

type UserTripDetailResponse struct {
	UserTrip db.UserTrip
}

type UserTripsDetailsResponse struct {
	UserTrips []db.UserTrip
}
