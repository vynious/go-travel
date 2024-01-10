-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
    LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    username,
    email
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: UpdateUserProfilePicture :one
UPDATE users
SET profile_picture = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateUserUsername :one
UPDATE users
SET username = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateUserName :one
UPDATE users
SET name = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
    RETURNING *;
