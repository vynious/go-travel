-- name: CreateMedia :one
INSERT INTO media (
                   entry_id,
                   key
) VALUES (
          $1, $2
         ) RETURNING *;


-- view media by entry id,

-- name: GetAllMediaByEntryId :many
SELECT *
FROM media
WHERE entry_id = $1;


-- name: GetAllMediaByTripId :many
SELECT media.*
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
WHERE travel_entry.trip_id = $1;



-- name: GetAllMediaByUserId :many
SELECT media.*
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
         JOIN user_trip ON travel_entry.trip_id = user_trip.trip_id
WHERE user_trip.user_id = $1;



-- name: GetAllMediaByTripIdAndUserId :many
SELECT media.*
FROM media
         JOIN travel_entry ON media.entry_id = travel_entry.id
         JOIN user_trip ON travel_entry.trip_id = user_trip.trip_id
WHERE user_trip.trip_id = $1 AND user_trip.user_id = $2;



-- name: GetMediaByKey :one
SELECT *
FROM media
WHERE entry_id = $1 AND key = $2;



-- name: UpdateMediaByKey :one
UPDATE media
SET key = $3
WHERE entry_id = $1 AND key = $2
RETURNING *;



-- name: DeleteMediaByKey :one
DELETE FROM media
WHERE entry_id = $1 AND key = $2
RETURNING *;


