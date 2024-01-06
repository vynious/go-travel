-- name: CreateConnection :one
INSERT INTO connections (
                         party_a,
                         party_b
) VALUES (
          $1, $2
         ) RETURNING *;

-- name: GetConnectionsByUserId :many
SELECT
    CASE
        WHEN party_a = $1 THEN party_b
        ELSE party_a
        END AS connected_user_id,
    connected_date
FROM connections
WHERE party_a = $1 OR party_b = $1;


-- name: DeleteConnectionByUserId :one
DELETE FROM connections
WHERE (party_a = $1 AND party_b = $2)
   OR (party_a = $2 AND party_b = $1)
RETURNING *;