package models

import (
	"time"

	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	NativeID     string
	ThreadID     string
	Snippet      string
	InternalDate time.Time
	Body         string
}
