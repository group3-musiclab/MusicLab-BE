package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	OrderID    string
	Status     string
	StudentID  uint
	MentorID   uint
	ClassID    uint
	ScheduleID uint
	StartDate  time.Time
	EndDate    time.Time
	Price      float64
	PaymentUrl string
	Student    Student
	Class      Class
	Duration   int
}

type Student struct {
	ID      uint
	Name    string
	Email   string
	Phone   string
	Address string
}

type Class struct {
	ID       uint
	Name     string
	Qty      int
	Price    float64
	Duration int
}
type TransactionHandler interface {
	MakeTransaction() echo.HandlerFunc
	GetMentorTransaction() echo.HandlerFunc
	GetStudentTransaction() echo.HandlerFunc
}

type TransactionService interface {
	MakeTransaction(newTransaction Core) (Core, error)
	GetMentorTransaction(mentorID uint) ([]Core, error)
	GetStudentTransaction(studentID uint) ([]Core, error)
}

type TransactionData interface {
	MakeTransaction(newTransaction Core) error
	GetMentorTransaction(mentorID uint) ([]Core, error)
	GetStudentTransaction(studentID uint) ([]Core, error)
}
