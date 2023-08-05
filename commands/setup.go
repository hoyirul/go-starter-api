package commands

import (
	"fmt"
	"log"

	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
	"github.com/hoyirul/go-starter-api/seeders"
)

func MigrateDB() {
	fmt.Println("Running mirgrate.....")
	// Inisialisasi koneksi ke database MySQL
	config.ConnectDB()

	// Migrate tabel user dan product
	config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	fmt.Println("migrate successfuly!")
}

func SeedDB() {
	config.ConnectDB()

	fmt.Println("Running seeders.....")
	// Menjalankan semua seeders
	if err := seeders.RunAll(); err != nil {
		log.Fatalf("Failed to run all seeders: %v", err)
	}

	fmt.Println("db:seed completed!")
}
