package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user", controllers.GetUserByID)
}
