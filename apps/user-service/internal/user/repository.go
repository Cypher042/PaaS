package user

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	FindUserByID(id uuid.UUID) (user *User)
	FindUserByUsername(username string) (user *User)
}
