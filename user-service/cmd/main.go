package main

import (
	"log"
	"os"

	"github.com/Cypher042/PaaS/user-service/internal/database"
	"github.com/Cypher042/PaaS/user-service/internal/repository"
	"github.com/Cypher042/PaaS/user-service/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}
	user.InitOAuth()
	database := database.Connect(os.Getenv("MONGODB_URI"))
	user_repo := &repository.UserRepo{
		Db: database,
	}
	user_service := user.NewService(user_repo)
	user_handler := user.NewHandler(user_service)
	r := gin.Default()
	user.RegisterRoutes(r, user_handler)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
