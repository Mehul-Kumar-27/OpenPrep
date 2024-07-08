package routes


import (
	"github.com/Mehul-Kumar-27/OpenBackend/auth"
)

import(
	"github.com/gin-gonic/gin"
)

type Routes interface {
	SetupRoutes() *gin.Engine

	HealthCheck(rg *gin.RouterGroup) *gin.RouterGroup
	AuthGroup(rg *gin.RouterGroup) *gin.RouterGroup
}


type AbstractRouter struct{
	engine *gin.Engine
	manualAuthService auth.AuthService
}


func NewAbstractRouter(engine *gin.Engine) *AbstractRouter{
	authController := auth.NewAuthController()
	return &AbstractRouter{
		engine: engine,
		manualAuthService: authController,
	}

}


func (router *AbstractRouter) SetupRoutes() *gin.Engine{

	api := router.engine.Group("/api/v1")
	router.AuthGroup(api)
	router.HealthCheck(api)
	return router.engine
}


func (ar *AbstractRouter) AuthGroup(rg *gin.RouterGroup) *gin.RouterGroup {
    authGroup := rg.Group("/auth")
    {
        authGroup.POST("/manual-login", ar.manualAuthService.Login)
        // Add other auth routes
    }
    return authGroup
}


func (router *AbstractRouter) HealthCheck(rg *gin.RouterGroup) *gin.RouterGroup{
	rg.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Health Check",
		})
	})
	return rg
}