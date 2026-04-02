// package router

// import (
// 	"github.com/Cypher042/PaaS/user-service/internal/user"
// 	"github.com/gin-gonic/gin"
// )
// func SetupRouter(h *user.Handler) *gin.Engine {

// 	r := gin.Default()

// 	r.GET("/auth/github", h.GithubLogin)
// 	r.GET("/auth/github/callback", h.GithubCallback)
// 	return r

// }


// //  RegisterRoutes for git module



// // func RegisterRoutes(r *gin.engine, h *handler) {
// // 	g := r.group("/user")

// // 	g.get("/auth/github", h.githublogin)
// // 	g.get("/auth/github/callback", h.githubcallback)
// // }
