package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Tag     string    `json:"tag"`
	PostTag []PostTag `gorm:"foreignKey:TagID"`
}
