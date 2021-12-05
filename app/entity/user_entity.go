package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null;unique" json:"name" form:"name" binding:"required"`
	Email    string `gorm:"size:255;not null;unique" json:"email" form:"email" binding:"required"`
	Password string `gorm:"size:255;not null;unique" json:"password" form:"password" binding:"required"`
}
