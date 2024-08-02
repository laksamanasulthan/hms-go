package user_models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey,autoIncrement" json:"id"`
	Email     string `json:"name"`
	Username  string `json:"price"`
	Password  string `json:"qty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
