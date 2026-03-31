package user

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	repo UserRepository
}

func NewService(r UserRepository) *Service {
	return &Service{repo: r}
}

func (s *Service) GithubCallback(code string) (*User, string, error) {
	token, err := GithubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, "", errors.New("Failed to exchange token")
	}

	user_data, err := GetGithubUser(token)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.FindUserByGithubUsername(user_data.GithubUsername)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			if err := s.repo.Create(user_data); err != nil {
				return nil, "", errors.New("Failed to create user")
			}
			return user_data, "", nil
		}
		return nil, "", err
	}
	// 1. Define your payload (claims)
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),                          // subject (the user ID)
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // expiration (7 days)
	}

	// 2. Create the token structural object
	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign the token using a secret key (from your .env)
	jwtSecret := []byte(os.Getenv("JWT_SECRET")) // Make sure you add JWT_SECRET to your .env

	jwtTokenStr, err := jwt_token.SignedString(jwtSecret)
	if err != nil {
		return nil, "", errors.New("Failed to generate token")
	}

	return user, jwtTokenStr, nil
}
