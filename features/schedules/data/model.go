package data

import (
	"musiclab-be/features/schedules"
	_modelTransaction "musiclab-be/features/transactions/data"
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	MentorID     uint
	Day          string `gorm:"type:enum('Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday')"`
	StartTime    string
	EndTime      string
	Transactions []_modelTransaction.Transaction
}

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

func ToCore(data Schedule) schedules.Core {
	return schedules.Core{
		ID:        data.ID,
		MentorID:  data.MentorID,
		Day:       data.Day,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	}
}

func CoreToData(data schedules.Core) Schedule {
	return Schedule{
		Model:     gorm.Model{ID: data.ID},
		MentorID:  data.MentorID,
		Day:       data.Day,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	}
}
