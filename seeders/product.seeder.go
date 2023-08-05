package seeders

import (
	"github.com/hoyirul/go-starter-api/config"
	"github.com/hoyirul/go-starter-api/models"
)

// ProductSeeder
type ProductSeeder struct{}

func (s ProductSeeder) Run() error {
	db := config.DB

	// Mengecek apakah tabel "products" sudah ada dalam basis data
	if !db.Migrator().HasTable(&models.Product{}) {
		return nil // Tabel belum ada, maka tidak perlu melakukan seeding
	}

	// Data produk yang akan di-seed
	products := []models.Product{
		{ID: 1, Name: "Product 1", Price: 1000},
		{ID: 2, Name: "Product 2", Price: 2000},
		// Tambahkan data produk lainnya sesuai kebutuhan
	}

	// Seed data produk ke dalam basis data
	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return err
		}
	}

	return nil
}
