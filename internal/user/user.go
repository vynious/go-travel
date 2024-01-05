package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	id        uuid.UUID
	name      string
	username  string
	email     string
	createdAt time.Time
}

type RegisterUserRes struct {
}

type LoginUserRes struct {
}
