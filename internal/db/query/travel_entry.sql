-- name: CreateTravelEntry :one
INSERT INTO travel_entry (
    user_id,
    trip_id,
    location,
    description
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetTravelEntryById :one
SELECT *
FROM travel_entry
WHERE id = $1
    LIMIT 1;

-- name: GetAllTravelEntry :many
SELECT *
FROM travel_entry;

-- name: GetAllTravelEntryByTripId :many
SELECT *
FROM travel_entry
WHERE trip_id = $1;

-- name: GetAllTravelEntryByUserIdAndTripId :many
SELECT *
FROM travel_entry
WHERE trip_id = $1
  AND user_id = $2;

-- name: UpdateTravelEntryLocation :one
UPDATE travel_entry
SET location = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateTravelEntryDescription :one
UPDATE travel_entry
SET description = $2
WHERE id = $1
    RETURNING *;


-- name: DeleteTravelEntry :one
DELETE FROM travel_entry
WHERE id = $1
    RETURNING *;
