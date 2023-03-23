package data

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	MentorID  uint
	StartDate time.Time
	EndDate   time.Time
}
