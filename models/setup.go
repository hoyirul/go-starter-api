package models

import (
	"time"

	"gorm.io/gorm"
)

type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Timestamps) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now().In(time.FixedZone("WIB", 7*60*60)) // Set zona waktu ke "Asia/Jakarta"
	t.CreatedAt = currentTime
	t.UpdatedAt = currentTime
	return
}

func (t *Timestamps) BeforeUpdate(tx *gorm.DB) (err error) {
	currentTime := time.Now().In(time.FixedZone("WIB", 7*60*60)) // Set zona waktu ke "Asia/Jakarta"
	t.UpdatedAt = currentTime
	return
}
