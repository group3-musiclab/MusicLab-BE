package data

import (
	"musiclab-be/features/transactions"
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
	PaymentUrl string
}

func ToCore(data Transaction) transactions.Core {
	return transactions.Core{
		ID:         data.ID,
		OrderID:    data.OrderID,
		Status:     data.Status,
		StudentID:  data.StudentID,
		MentorID:   data.MentorID,
		ClassID:    data.ClassID,
		ScheduleID: data.ScheduleID,
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
		Price:      data.Price,
		PaymentUrl: data.PaymentUrl,
	}
}

func CoreToData(data transactions.Core) Transaction {
	return Transaction{
		Model:      gorm.Model{ID: data.ID},
		OrderID:    data.OrderID,
		Status:     data.Status,
		StudentID:  data.StudentID,
		MentorID:   data.MentorID,
		ClassID:    data.ClassID,
		ScheduleID: data.ScheduleID,
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
		Price:      data.Price,
		PaymentUrl: data.PaymentUrl,
	}
}
