package entity

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title  string `gorm:"size:255;not null;" json:"title" form:"title" binding:"required"`
	Text   string `gorm:"text;not null;" json:"text" form:"text" binding:"required"`
	User   User
	UserID uint32 `gorm:"not null" json:"user_id"`
}
