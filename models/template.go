package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	Name    string
	Subject string
	Body    string
	Sender  string
	UserID  uint

	Campaigns []Campaign `gorm:"many2many:template_campaigns;"`
	User      User       `gorm:"foreignKey:UserID;references:ID"`
}
