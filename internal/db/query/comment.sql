-- name: CreateComment :one
INSERT INTO comment (
                     entry_id,
                     user_id,
                     content
) VALUES (
          $1, $2, $3
         ) RETURNING *;


-- name: DeleteCommentById :one
DELETE FROM comment
WHERE id = $1
RETURNING *;