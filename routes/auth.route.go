package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
)

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", controllers.LoginHandler)
	}
}
