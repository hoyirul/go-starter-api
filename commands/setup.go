package commands

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

// GenerateSecretKey menghasilkan secret key secara acak dan menyimpannya di file .env
func GenerateSecretKey() {
	// Baca isi file .env
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		fmt.Println("Failed to read .env file:", err)
		return
	}

	// Konversi isi file .env menjadi string
	envContent := string(data)

	// Cek apakah SECRET_KEY sudah ada di dalam file .env
	if strings.Contains(envContent, "SECRET_KEY=") {
		// Jika SECRET_KEY sudah ada, ambil nilai lama
		oldSecretKey := getValueFromEnv(envContent, "SECRET_KEY")

		// Tambahkan pesan konfirmasi untuk mengganti secret key
		fmt.Println("SECRET_KEY already exists in .env. Do you want to generate a new one? (y/n)")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToLower(answer) == "n" {
			fmt.Println("Secret key generation aborted.")
			return
		}

		// Hapus nilai lama dari file .env
		envContent = strings.Replace(envContent, fmt.Sprintf("\nSECRET_KEY=%s", oldSecretKey), "", 1)
	}

	// Buat buffer untuk menampung secret key yang dihasilkan
	secretKey := make([]byte, 32)

	// Baca secara acak dari crypto/rand ke buffer secretKey
	_, err = rand.Read(secretKey)
	if err != nil {
		fmt.Println("Failed to generate secret key:", err)
		return
	}

	// Encode secret key ke dalam bentuk base64
	encodedSecretKey := base64.URLEncoding.EncodeToString(secretKey)

	// Simpan secret key ke dalam file .env
	if err := updateEnv("\nSECRET_KEY", encodedSecretKey, envContent); err != nil {
		fmt.Println("Failed to update SECRET_KEY in .env:", err)
		return
	}

	fmt.Println("Secret key generated and updated in .env successfully.")
}

func getValueFromEnv(envContent, key string) string {
	startIndex := strings.Index(envContent, key)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(key) + 1 // start index setelah "="
	endIndex := startIndex
	for ; endIndex < len(envContent); endIndex++ {
		if envContent[endIndex] == '\n' {
			break
		}
	}
	return envContent[startIndex:endIndex]
}

func updateEnv(key, value, envContent string) error {
	// Buka file .env untuk ditulis
	file, err := os.OpenFile(".env", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Tambahkan atau perbarui nilai variabel di .env
	_, err = fmt.Fprint(file, envContent)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(file, "%s=%s\n", key, value)
	return err
}
