package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
)

// GetAllProducts mengambil semua produk dari basis data dan mengirimkan sebagai JSON response
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID mengambil satu produk berdasarkan ID dari basis data dan mengirimkan sebagai JSON response
func GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, productID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct membuat produk baru berdasarkan data yang diberikan dan mengirimkan sebagai JSON response
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct mengupdate data produk berdasarkan ID yang diberikan dan mengirimkan sebagai JSON response
func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var product models.Product
	result := config.DB.First(&product, productID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price

	config.DB.Save(&product)

	c.JSON(http.StatusOK, product)
}

// DeleteProduct menghapus data produk berdasarkan ID yang diberikan dan mengirimkan sebagai JSON response
func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	result := config.DB.First(&product, productID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
