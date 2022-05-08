package main

import (
	"fmt"
	"log"

	"backend/config"
	helpers "backend/helpers"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

var (
	ginConfig config.GinConfig
	router    *gin.Engine
)

func initGinServer() {
	ginConfig.SERVER_NAME = helpers.GetEnv("SERVER_NAME", "ERP_BACKEND")
	ginConfig.SERVER_PORT = helpers.GetEnv("SERVER_PORT", "9090")
	ginConfig.SERVER_ENV = helpers.GetEnv("SERVER_ENV", "dev")

	config.SetGinMode(ginConfig.SERVER_ENV)
	router = routes.SetupRouter()
}

func main() {
	initGinServer()

	authService := helpers.GetEnv("AUTH_SERVICE", "http://localhost:8081")

	config.App = config.NewServices(router, &ginConfig, ginConfig.SERVER_ENV, authService)

	log.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%s", ginConfig.SERVER_PORT)))
}
