// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID          int64     `json:"id"`
	EntryID     int64     `json:"entry_id"`
	UserID      string    `json:"user_id"`
	Content     string    `json:"content"`
	CommentedOn time.Time `json:"commented_on"`
}

type Connection struct {
	PartyA        string    `json:"party_a"`
	PartyB        string    `json:"party_b"`
	ConnectedDate time.Time `json:"connected_date"`
}

type Medium struct {
	EntryID int64  `json:"entry_id"`
	Key     string `json:"key"`
}

type TravelEntry struct {
	ID          int64  `json:"id"`
	UserID      string `json:"user_id"`
	TripID      int64  `json:"trip_id"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type Trip struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Country   string    `json:"country"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type User struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	ProfilePicture sql.NullString `json:"profile_picture"`
	CreationDate   time.Time      `json:"creation_date"`
}

type UserTrip struct {
	TripID int64  `json:"trip_id"`
	UserID string `json:"user_id"`
}
