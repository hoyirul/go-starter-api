package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
	"github.com/hoyirul/go-starter-api/middlewares"
)

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", controllers.LoginHandler)
		authGroup.POST("/register", controllers.RegisterHandler)
	}

	router.Use(middlewares.JWTMiddleware())
	authGroup.POST("/logout", controllers.LogoutHandler)
}
