package user

import (
	"github.com/google/uuid"
)


type GithubResponse struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
}

type User struct {
	ID       uuid.UUID `bson:"id"`
	Username string    `bson:"username"`
	GithubUsername string `bson:"githubusername"`
	GithubToken string `bson:"github_token"`
	Email string `bson:"email"`
}

// DTOs

