package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/commands"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/middlewares"
	"github.com/hoyirul/go-starter-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [function]")
		return
	}

	functionToRun := os.Args[1]

	switch functionToRun {
	case "key:generate":
		commands.GenerateSecretKey()
	case "db:seed":
		commands.SeedDB()
	case "migrate":
		commands.MigrateDB()
	case "serve":
		config.ConnectDB()
		if key := os.Getenv("SECRET_KEY"); key == "" {
			fmt.Println("Your key is null, please run `go run main.go key:generate`")
		} else {
			fmt.Println("Starting Go API Server...")

			// Baca variabel lingkungan dari file .env
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			// Periksa apakah variabel GIN_MODE ada di lingkungan
			ginMode := os.Getenv("GIN_MODE")
			if ginMode == "development" {
				ginMode = gin.DebugMode
			}

			// Set mode Gin
			gin.SetMode(ginMode)

			// Inisialisasi router Gin
			// router := gin.Default()
			router := gin.New()

			// Untuk CORS
			router.Use(middlewares.CorsMiddleware())

			// Setup rute untuk pengguna (user) dan produk (product)
			routes.SetupRoutes(router)

			// Jalankan server
			router.Run(":8080")
		}
	default:
		fmt.Println("Function not found.")
	}
}
