package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GithubLogin(c *gin.Context) {
	state := "random"
	url := GithubOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handler) GithubCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	if state != "random" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code is missing"})
		return
	}

	_, jwtToken, err := h.service.GithubCallback(code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", jwtToken, 60*60*24*7, "/", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/dashboard")
}
