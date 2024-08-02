package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string    `json:"title"`
	Post    string    `json:"post"`
	View    int       `json:"view"`
	PostTag []PostTag `gorm:"foreignKey:TagID"`
}
