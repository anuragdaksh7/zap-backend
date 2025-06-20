package models

import (
	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	UserID             uint
	GoogleAccessToken  string
	GoogleRefreshToken string

	User User `gorm:"foreignKey:UserID;references:ID"`
}
