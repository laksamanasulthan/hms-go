package models

import "time"

type Product struct {
	ID        uint   `gorm:"primarykey,autoIncrement" json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Qty       uint   `json:"qty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
