package models

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Timestamps
}

func (User) TableName() string {
	return "users"
}
