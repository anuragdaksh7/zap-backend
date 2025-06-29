package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	UserID             uint
	Name               string
	Description        string
	RunEveryNMinutes   uint
	TotalProspects     uint
	ProcessedProspects uint `gorm:"default:0"`
	ProcessedAt        *time.Time
	CurrentStatus      string

	User      User       `gorm:"foreignKey:UserID;references:ID"`
	Templates []Template `gorm:"many2many:template_campaigns;"`
}
