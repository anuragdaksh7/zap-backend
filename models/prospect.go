package models

import (
	"time"

	"gorm.io/gorm"
)

type Prospect struct {
	gorm.Model
	CampaignID  uint
	FullName    string
	Email       string
	Phone       string
	Company     string
	Position    string
	LinkedInURL string
	Status      string
	Location    string
	LastContact time.Time
	ResponseAt  time.Time
	Notes       string

	Campaign Campaign `gorm:"foreignKey:CampaignID;references:ID"`
}
