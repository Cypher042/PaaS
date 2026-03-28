package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo UserRepository
}

func NewService(r UserRepository) *Service {
	return &Service{repo: r}
}


func (s *Service) Register(regRequest RegisterRequest) (*User, error){

	username := regRequest.Username

	if(s.repo.FindUserByUsername(username) != nil){
		return nil, fmt.Errorf("A user with the username %v already exists", username)
	}

	

	user = s.repo.Create()
}