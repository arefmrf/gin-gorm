package models

import (
	"gorm.io/gorm"
	"web/internal/modules/user/models"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"varchar(191)"`
	Content  string `gorm:"text"`
	Password string `gorm:"varchar(191)"`
	UserID   uint
	User     models.User
}
