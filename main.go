package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
	"github.com/hoyirul/go-starter-api/routes"
)

func main() {
	fmt.Println("Starting Go API Server...")

	// Inisialisasi koneksi ke database MySQL
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate tabel user dan product
	config.DB.AutoMigrate(&models.User{}, &models.Product{})

	// Inisialisasi router Gin
	router := gin.Default()

	// Setup rute untuk pengguna (user) dan produk (product)
	routes.SetupUserRoutes(router)
	routes.SetupProductRoutes(router)

	// Jalankan server
	router.Run(":8080")
}
