package router

import (
	"github.com/Cypher042/PaaS/user-service/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *user.Handler) *gin.Engine {

	r := gin.Default()

	r.GET("/auth/github", h.GithubLogin)
	r.GET("/auth/github/callback", h.GithubCallback)
	return r

}
