package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/controllers"
)

func SetupProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/api/products")
	{
		productGroup.GET("", controllers.GetAllProducts)
		productGroup.GET("/:id", controllers.GetProductByID)
		productGroup.POST("", controllers.CreateProduct)
		productGroup.PUT("/:id", controllers.UpdateProduct)
		productGroup.DELETE("/:id", controllers.DeleteProduct)
	}
}
