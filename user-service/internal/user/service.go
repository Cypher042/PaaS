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

	if s.repo.FindUserByUsername(username) != nil {
		return nil, fmt.Errorf("A user with the username %v already exists", username)
	}

	user := User{}
	user.ID = uuid.New()
	user.Username = regRequest.Username
	user.Password = regRequest.Password
	user.Github = regRequest.Github

	err := s.repo.Create(&user)

	if err != nil {
		return nil, errors.New()
	}

	return &user, nil

}

func (s *Service) Login(loginRequest LoginRequest) (*User, error){

	username := loginRequest.Username
	
	user := s.repo.FindUserByUsername(username)
	if user  == nil {
		return nil, fmt.Errorf("A user with the username %v does not exist", username)
	}

	if err != nil {
		return nil, Error.New()
	}

	return &user, nil

}