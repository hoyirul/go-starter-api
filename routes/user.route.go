package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.POST("", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
