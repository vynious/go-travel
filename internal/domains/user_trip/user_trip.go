package user_trip

import db "github.com/vynious/go-travel/internal/db/sqlc"

type AddUsersTripRequest struct {
	Users []string
}

type UserTripDetailResponse struct {
	UserTrip db.UserTrip
}

type MultipleUserTripDetailResponse struct {
	UserTrips []db.UserTrip
}

type UserTripsDetailsResponse struct {
	UserTrips []db.UserTrip
}
