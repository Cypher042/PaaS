package router

import (
	"github.com/Cypher042/PaaS/user-service/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *user.Handler) *gin.Engine {

	r := gin.Default()

	r.GET("/login", h.Login)
	r.POST("/register", h.Register)
	return r

}
