package data

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID   string
	Status    string
	StudentID uint
	MentorID  uint
	ClassID   uint
	StartDate time.Time
	EndDate   time.Time
	Price     float64 `gorm:"type:float"`
}
