package internal


import (

	"github.com/google/uuid"

)

type UserRepository interface{

	Create(user *User) (error)
	FindUserByID(id uuid.UUID) (user *User)
}
