package data

import (
	_modelTransaction "musiclab-be/features/transactions/data"
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	MentorID     uint
	Day          string    `gorm:"type:enum('Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday')"`
	StartTime    time.Time `gorm:"type:time"`
	EndTime      time.Time `gorm:"type:time"`
	Transactions []_modelTransaction.Transaction
}
