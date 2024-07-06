package main

import (
	"time"

	"github.com/Mehul-Kumar-27/OpenBackend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Mehul-Kumar-27/OpenBackend/routes"
)

func main() {

	config.Init()

	config.InitLogger(config.AppConfig.LogLevel, config.AppConfig.LogFile, config.AppConfig.LogToConsole)

	router := gin.New()

	router.Use(
		cors.New(
			cors.Config{
				AllowAllOrigins: true,
				AllowMethods:     []string{"OPTIONS", "PUT", "PATCH", "GET", "DELETE"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Origin"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		),
	)

	router.Use(gin.Recovery())

	router.Use(config.LogMiddleware())

	abstractRouter := routes.NewAbstractRouter(router)

	
	abstractRouter.SetupRoutes()

	router.Run(":" + config.AppConfig.ServerPort)
	
}
