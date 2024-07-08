package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Mehul-Kumar-27/OpenBackend/config"
	"github.com/Mehul-Kumar-27/OpenBackend/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Mehul-Kumar-27/OpenBackend/routes"
)

func main() {
	config.Init()
	fmt.Print("Starting server...")
	config.InitLogger(config.AppConfig.LogLevel, config.AppConfig.LogFile, config.AppConfig.LogToConsole)
	config.Log.Info("Logger initialized")
	database.InitializeDBHandler()
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
	
	config.Log.Infof("Starting server on port %s", config.AppConfig.ServerPort)
	err := router.Run(":" + config.AppConfig.ServerPort)
	if err != nil {
		config.Log.Fatalf("Error starting server: %s", err.Error())
		os.Exit(1)
	}
	
}
