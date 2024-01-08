package user

import (
	db "github.com/vynious/go-travel/internal/db/sqlc"
)

type RegisterUserRequest struct {
	User     db.User
	Password string
}

type UserDetailResponse struct {
	User db.User
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type UpdateUserDetailRequest struct {
	Name     *string
	Email    *string
	Username *string
}
