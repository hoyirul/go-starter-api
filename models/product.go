package models

type Product struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Timestamps
}

func (Product) TableName() string {
	return "products"
}
