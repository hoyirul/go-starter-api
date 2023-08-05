package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	SetupAuthRoutes(router)

	router.Use(middlewares.JWTMiddleware())
	SetupProductRoutes(router)
	SetupUserRoutes(router)
}
