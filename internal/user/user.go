package user

import (
	db "github.com/vynious/go-travel/internal/db/sqlc"
)

type RegisterUserReq struct {
	User     db.User
	Password string
}

type RegisterUserRes struct {
	User db.User
}

type LoginUserRes struct {
	Email    string
	Password string
}
