package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	IsVerified bool
	OTP        string
	ProfilePic string

	UserToken *UserToken `gorm:"foreignKey:UserID;references:ID"`
	Campaign  []Campaign `gorm:"foreignKey:UserID;references:ID"`
}
