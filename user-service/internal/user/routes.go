package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *Handler) {
	g := r.Group("/user")

	g.GET("/auth/github", h.GithubLogin)
	g.GET("/auth/github/callback", h.GithubCallback)
}
