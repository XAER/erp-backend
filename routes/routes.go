package routes

import (
	"backend/controllers"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.PingHandler)
		v1.POST("/register", controllers.RegisterHandler)
		v1.POST("/login", controllers.LoginHandler)
		// From here, there will be only authenticated routes
		// TODO implement checkToken middleware
		v1.GET("/pingWithToken", middleware.AuthorizationMiddleware(), middleware.TokenAuthMiddleware(), controllers.PingHandlerWithToken)
	}

	return r
}
