package router

import (
	"github.com/Cypher042/PaaS/user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/login", handler.Login)
	r.POST("/register", handler.Login)
	return r

}

