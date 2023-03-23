package data

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID    string
	Status     string
	StudentID  uint
	MentorID   uint
	ClassID    uint
	ScheduleID uint
	StartDate  time.Time `gorm:"type:date"`
	EndDate    time.Time `gorm:"type:date"`
	Price      float64   `gorm:"type:float"`
}
