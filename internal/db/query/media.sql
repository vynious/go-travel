-- name: CreateMedia :one
INSERT INTO media (
                   entry_id,
                   url
) VALUES (
          $1, $2
         ) RETURNING *;




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




-- name: UpdateMediaById :one
UPDATE media
SET url = $2
WHERE id = $1
    RETURNING *;



-- name: DeleteMediaById :one
DELETE FROM media
WHERE id = $1
    RETURNING *;


