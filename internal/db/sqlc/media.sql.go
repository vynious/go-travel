// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: media.sql

package db

import (
	"context"
)

const createMedia = `-- name: CreateMedia :one
INSERT INTO media (
                   entry_id,
                   key
) VALUES (
          $1, $2
         ) RETURNING entry_id, key
`

type CreateMediaParams struct {
	EntryID int64  `json:"entry_id"`
	Key     string `json:"key"`
}

func (q *Queries) CreateMedia(ctx context.Context, arg CreateMediaParams) (Medium, error) {
	row := q.queryRow(ctx, q.createMediaStmt, createMedia, arg.EntryID, arg.Key)
	var i Medium
	err := row.Scan(&i.EntryID, &i.Key)
	return i, err
}

const deleteMediaByKey = `-- name: DeleteMediaByKey :one
DELETE FROM media
WHERE entry_id = $1 AND key = $2
RETURNING entry_id, key
`

type DeleteMediaByKeyParams struct {
	EntryID int64  `json:"entry_id"`
	Key     string `json:"key"`
}

func (q *Queries) DeleteMediaByKey(ctx context.Context, arg DeleteMediaByKeyParams) (Medium, error) {
	row := q.queryRow(ctx, q.deleteMediaByKeyStmt, deleteMediaByKey, arg.EntryID, arg.Key)
	var i Medium
	err := row.Scan(&i.EntryID, &i.Key)
	return i, err
}

const getAllMediaByEntryId = `-- name: GetAllMediaByEntryId :many

SELECT entry_id, key
FROM media
WHERE entry_id = $1
`

// view media by entry id,
func (q *Queries) GetAllMediaByEntryId(ctx context.Context, entryID int64) ([]Medium, error) {
	rows, err := q.query(ctx, q.getAllMediaByEntryIdStmt, getAllMediaByEntryId, entryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Medium{}
	for rows.Next() {
		var i Medium
		if err := rows.Scan(&i.EntryID, &i.Key); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMediaByTripId = `-- name: GetAllMediaByTripId :many
SELECT media.entry_id, media.key
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
WHERE travel_entry.trip_id = $1
`

func (q *Queries) GetAllMediaByTripId(ctx context.Context, tripID int64) ([]Medium, error) {
	rows, err := q.query(ctx, q.getAllMediaByTripIdStmt, getAllMediaByTripId, tripID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Medium{}
	for rows.Next() {
		var i Medium
		if err := rows.Scan(&i.EntryID, &i.Key); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMediaByTripIdAndUserId = `-- name: GetAllMediaByTripIdAndUserId :many
SELECT media.entry_id, media.key
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
         JOIN user_trip ON travel_entry.trip_id = user_trip.trip_id
WHERE user_trip.trip_id = $1 AND user_trip.user_id = $2
`

type GetAllMediaByTripIdAndUserIdParams struct {
	TripID int64  `json:"trip_id"`
	UserID string `json:"user_id"`
}

func (q *Queries) GetAllMediaByTripIdAndUserId(ctx context.Context, arg GetAllMediaByTripIdAndUserIdParams) ([]Medium, error) {
	rows, err := q.query(ctx, q.getAllMediaByTripIdAndUserIdStmt, getAllMediaByTripIdAndUserId, arg.TripID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Medium{}
	for rows.Next() {
		var i Medium
		if err := rows.Scan(&i.EntryID, &i.Key); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMediaByUserId = `-- name: GetAllMediaByUserId :many
SELECT media.entry_id, media.key
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
         JOIN user_trip ON travel_entry.trip_id = user_trip.trip_id
WHERE user_trip.user_id = $1
`

func (q *Queries) GetAllMediaByUserId(ctx context.Context, userID string) ([]Medium, error) {
	rows, err := q.query(ctx, q.getAllMediaByUserIdStmt, getAllMediaByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Medium{}
	for rows.Next() {
		var i Medium
		if err := rows.Scan(&i.EntryID, &i.Key); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMediaByKey = `-- name: GetMediaByKey :one
SELECT entry_id, key
FROM media
WHERE entry_id = $1 AND key = $2
`

type GetMediaByKeyParams struct {
	EntryID int64  `json:"entry_id"`
	Key     string `json:"key"`
}

func (q *Queries) GetMediaByKey(ctx context.Context, arg GetMediaByKeyParams) (Medium, error) {
	row := q.queryRow(ctx, q.getMediaByKeyStmt, getMediaByKey, arg.EntryID, arg.Key)
	var i Medium
	err := row.Scan(&i.EntryID, &i.Key)
	return i, err
}

const updateMediaByKey = `-- name: UpdateMediaByKey :one
UPDATE media
SET key = $3
WHERE entry_id = $1 AND key = $2
RETURNING entry_id, key
`

type UpdateMediaByKeyParams struct {
	EntryID int64  `json:"entry_id"`
	Key     string `json:"key"`
	Key_2   string `json:"key_2"`
}

func (q *Queries) UpdateMediaByKey(ctx context.Context, arg UpdateMediaByKeyParams) (Medium, error) {
	row := q.queryRow(ctx, q.updateMediaByKeyStmt, updateMediaByKey, arg.EntryID, arg.Key, arg.Key_2)
	var i Medium
	err := row.Scan(&i.EntryID, &i.Key)
	return i, err
}
