package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Timestamps
}

func (User) TableName() string {
	return "users"
}

// HashPassword mengenkripsi password pengguna menggunakan bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// BeforeCreate adalah hook yang akan dijalankan sebelum data pengguna baru disimpan ke database
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.Password) > 0 {
		if err = u.HashPassword(); err != nil {
			return err
		}
	}
	return nil
}

// CheckPassword memeriksa apakah password yang diberikan cocok dengan password yang di-hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
