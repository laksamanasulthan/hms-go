package model

import (
	"gorm.io/gorm"
)

type PostTag struct {
	gorm.Model
	PostID uint `json:"post_id"`
	TagID  uint `json:"tag_id"`
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tag    Tag  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
