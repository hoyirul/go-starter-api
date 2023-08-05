package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	SetupProductRoutes(router)
	SetupUserRoutes(router)
}
