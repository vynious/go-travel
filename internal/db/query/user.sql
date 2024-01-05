-- name: GetUser :one
select * from "user"
where id = $1
limit 1;

-- name: ListUsers :many
select * from "user"
order by name;

-- name: CreateUser :one
insert into "user" (
    name,
    username,
    email
) values (
    $1, $2, $3
) returning *;

-- name: GetUserByEmail :one
select email from "user"
where username = $1;

-- name: UpdateUserProfilePicture :one
update user
set profile_picture = $2
where id = $1;
