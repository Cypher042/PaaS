package user

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *User) error
	FindUserByID(id uuid.UUID) (user *User, err error)
	FindUserByUsername(username string) (user *User, err error)
	
}
