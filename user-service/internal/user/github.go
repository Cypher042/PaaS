package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOAuthConfig *oauth2.Config

func InitOAuth() {
	GithubOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/github/callback",
		// "repo" scope allows access to private repositories.
		// "user:email" gets their email address.
		// "user:username" gets their username.
		Scopes: []string{"user:email", "repo"},
	}
}

func GetGithubUser(token *oauth2.Token) (*User, error) {
	client := GithubOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, errors.New("Failed to get user")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get user")
	}

	var github_resp GithubResponse
	if err := json.NewDecoder(resp.Body).Decode(&github_resp); err != nil {
		return nil, errors.New("Failed to decode user")
	}

	if github_resp.Login == "" {
		return nil, errors.New("Github api returned an empty username")
	}

	user_data_resp := &User{
		ID:             uuid.New(),
		Username:       github_resp.Login,
		GithubUsername: github_resp.Login,
		Email:          github_resp.Email,
		GithubToken:    token.AccessToken,
	}
	return user_data_resp, nil
}
