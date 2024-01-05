
-- name: GetTravelEntriesByOwnerId :many
select country, street_address, city, state, postal_code, media, description, visit_date
from travel_entry
where owner_id = $1;

-- name: GetTravelEntriesById :one
select country, street_address, city, state, postal_code, media, description, visit_date
from travel_entry
where id = $1;

-- name: GetTravelEntriesByCountry :many
select country, street_address, city, state, postal_code, media, description, visit_date
from travel_entry
where country = $1;

-- name: CreateTravelEntry :one
insert into travel_entry (
    owner_id,
    country,
    street_address,
    city,
    state,
    postal_code,
    media,
    description,
    visibility,
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) returning *;

