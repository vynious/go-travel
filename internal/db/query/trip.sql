-- name: CreateTrip :one
INSERT INTO trip (
    title,
    country,
    start_date,
    end_date
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetTrip :one
SELECT * FROM trip
WHERE id = $1;

-- name: ListTrips :many
SELECT * FROM trip
ORDER BY start_date;

-- name: UpdateTripStartDate :one
UPDATE trip
SET start_date = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateTripEndDate :one
UPDATE trip
SET end_date = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateTripCountry :one
UPDATE trip
SET country = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateTripTitle :one
UPDATE trip
SET title = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteTrip :one
DELETE FROM trip
WHERE id = $1
    RETURNING *;
