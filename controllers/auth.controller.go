package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/middlewares"
	"github.com/hoyirul/go-starter-api/models"
)

// LoginRequest adalah struktur untuk menerima data login dari permintaan
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse adalah struktur untuk menanggapi hasil login
type LoginResponse struct {
	TokenType string      `json:"token_type"`
	Token     string      `json:"token"`
	User      models.User `json:"user"`
}

// LoginHandler menghandle proses login dan menghasilkan token JWT beserta data user
func LoginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Cek apakah username ada di database
	var user models.User
	if err := config.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Memeriksa password yang diberikan dengan password yang di-hash
	if err := user.CheckPassword(request.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Jika login berhasil, hasilkan token JWT
	token, err := middlewares.GenerateJWTToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token and user data in the response
	c.JSON(http.StatusOK, LoginResponse{
		TokenType: "Bearer",
		Token:     token,
		User:      user,
	})
}

// RegisterRequest adalah struktur untuk menerima data input saat registrasi
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterResponse adalah struktur untuk menanggapi hasil registrasi
type RegisterResponse struct {
	TokenType string      `json:"token_type"`
	Token     string      `json:"token"`
	User      models.User `json:"user"`
}

// RegisterHandler menghandle proses registrasi dan menghasilkan token JWT beserta data user
func RegisterHandler(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Cek apakah username sudah terdaftar
	var existingUser models.User
	if err := config.DB.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	token, err := middlewares.GenerateJWTToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token and user data in the response
	c.JSON(http.StatusCreated, RegisterResponse{
		TokenType: "Bearer",
		Token:     token,
		User:      user,
	})
}

// LogoutHandler menghandle proses logout dan menghapus token dari header Authorization
func LogoutHandler(c *gin.Context) {
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
