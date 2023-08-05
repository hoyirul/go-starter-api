package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
)

// GetAllUsers mengambil semua pengguna dari basis data dan mengirimkan sebagai JSON response
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID mengambil satu pengguna berdasarkan ID dari basis data dan mengirimkan sebagai JSON response
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser membuat pengguna baru berdasarkan data yang diberikan dan mengirimkan sebagai JSON response
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser mengupdate data pengguna berdasarkan ID yang diberikan dan mengirimkan sebagai JSON response
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	config.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

// DeleteUser menghapus data pengguna berdasarkan ID yang diberikan dan mengirimkan sebagai JSON response
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := config.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
