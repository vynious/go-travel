// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateConnection(ctx context.Context, arg CreateConnectionParams) (Connection, error)
	CreateMedia(ctx context.Context, arg CreateMediaParams) (Medium, error)
	CreateTravelEntry(ctx context.Context, arg CreateTravelEntryParams) (TravelEntry, error)
	CreateTrip(ctx context.Context, arg CreateTripParams) (Trip, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCommentById(ctx context.Context, id int64) (Comment, error)
	DeleteConnectionByUserId(ctx context.Context, arg DeleteConnectionByUserIdParams) (Connection, error)
	DeleteMediaById(ctx context.Context, id int64) (Medium, error)
	DeleteTravelEntry(ctx context.Context, id int64) (TravelEntry, error)
	DeleteTrip(ctx context.Context, id int64) (Trip, error)
	DeleteUser(ctx context.Context, id string) (User, error)
	GetAllMediaByEntryId(ctx context.Context, entryID sql.NullInt64) ([]Medium, error)
	GetAllMediaByTripId(ctx context.Context, tripID sql.NullInt64) ([]Medium, error)
	GetAllMediaByTripIdAndUserId(ctx context.Context, arg GetAllMediaByTripIdAndUserIdParams) ([]Medium, error)
	GetAllMediaByUserId(ctx context.Context, userID string) ([]Medium, error)
	GetAllTravelEntry(ctx context.Context) ([]TravelEntry, error)
	GetAllTravelEntryByTripId(ctx context.Context, tripID sql.NullInt64) ([]TravelEntry, error)
	GetAllTravelEntryByUserIdAndTripId(ctx context.Context, arg GetAllTravelEntryByUserIdAndTripIdParams) ([]TravelEntry, error)
	GetConnectionsByUserId(ctx context.Context, partyA string) ([]GetConnectionsByUserIdRow, error)
	GetTravelEntryById(ctx context.Context, id int64) (TravelEntry, error)
	GetTrip(ctx context.Context, id int64) (Trip, error)
	GetUser(ctx context.Context, id string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListTrips(ctx context.Context) ([]Trip, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateMediaById(ctx context.Context, arg UpdateMediaByIdParams) (Medium, error)
	UpdateTravelEntryDescription(ctx context.Context, arg UpdateTravelEntryDescriptionParams) (TravelEntry, error)
	UpdateTravelEntryLocation(ctx context.Context, arg UpdateTravelEntryLocationParams) (TravelEntry, error)
	UpdateTripCountry(ctx context.Context, arg UpdateTripCountryParams) (Trip, error)
	UpdateTripStartDate(ctx context.Context, arg UpdateTripStartDateParams) (Trip, error)
	UpdateTripTitle(ctx context.Context, arg UpdateTripTitleParams) (Trip, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error)
	UpdateUserProfilePicture(ctx context.Context, arg UpdateUserProfilePictureParams) (User, error)
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (User, error)
}

var _ Querier = (*Queries)(nil)
