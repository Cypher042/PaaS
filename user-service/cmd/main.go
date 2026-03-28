package main

import (
	"github.com/Cypher042/PaaS/user-service/internal/router"
)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
