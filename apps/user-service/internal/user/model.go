package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `bson:"id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
	Github   string    `bson:"github"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Github   string `json:"github"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Github   string `json:"github"`
}
