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
	Student    Student
	Class      Class
	Mentor     Mentor
}

type Student struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Address string
}

type Mentor struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Address string
}

type Class struct {
	gorm.Model
	Name     string
	Duration int
	Price    float64
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
		Student: transactions.Student{
			ID:      data.Student.ID,
			Name:    data.Student.Name,
			Email:   data.Student.Email,
			Phone:   data.Student.Phone,
			Address: data.Student.Address,
		},
		Class: transactions.Class{
			ID:       data.Class.ID,
			Name:     data.Class.Name,
			Duration: data.Class.Duration,
			Price:    data.Class.Price,
		},
		Mentor: transactions.Mentor{
			ID:      data.Mentor.ID,
			Name:    data.Mentor.Name,
			Email:   data.Mentor.Email,
			Phone:   data.Mentor.Phone,
			Address: data.Mentor.Address,
		},
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
