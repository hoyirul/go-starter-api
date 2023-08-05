package seeders

import (
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
)

// UserSeeder
type UserSeeder struct{}

func (s UserSeeder) Run() error {
	db := config.DB

	// Mengecek apakah tabel "Users" sudah ada dalam basis data
	if !db.Migrator().HasTable(&models.User{}) {
		return nil // Tabel belum ada, maka tidak perlu melakukan seeding
	}

	// Data produk yang akan di-seed
	users := []models.User{
		{Username: "admin", Email: "admin@gmail.com"},
		{Username: "user", Email: "user@gmail.com"},
		// Tambahkan data produk lainnya sesuai kebutuhan
	}

	// Seed data produk ke dalam basis data
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
