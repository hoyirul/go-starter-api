package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-api/commands"
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/routes"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [function]")
		return
	}

	functionToRun := os.Args[1]

	switch functionToRun {
	case "db:seed":
		commands.SeedDB()
	case "migrate":
		commands.MigrateDB()
	case "serve":
		config.ConnectDB()
		fmt.Println("Starting Go API Server...")
		// Inisialisasi router Gin
		// router := gin.Default()
		router := gin.New()

		// Setup rute untuk pengguna (user) dan produk (product)
		routes.SetupRoutes(router)

		// Jalankan server
		router.Run(":8080")
	default:
		fmt.Println("Function not found.")
	}
}
