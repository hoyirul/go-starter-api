package seeders

type Seeder interface {
	Run() error
}

func RunAll() error {
	// Panggil semua seeder yang ingin dijalankan di sini
	seeders := []Seeder{
		ProductSeeder{},
		UserSeeder{},
		// Tambahkan seeder lainnya sesuai kebutuhan
	}

	for _, seeder := range seeders {
		if err := seeder.Run(); err != nil {
			return err
		}
	}

	return nil
}
