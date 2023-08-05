package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
)

func SetupProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/product", controllers.GetProductByID)
}
