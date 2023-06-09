package schedules

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	MentorID    uint
	Day         string
	StartTime   string
	EndTime     string
	ClassID     uint `validate:"required"`
	Transaction Transaction
}

type Transaction struct {
	ID         uint
	OrderID    string
	Status     string
	StudentID  uint
	MentorID   uint
	ClassID    uint
	ScheduleID uint      `validate:"required"`
	StartDate  time.Time `validate:"required"`
	EndDate    time.Time
	Price      float64
}

type ScheduleHandler interface {
	PostSchedule() echo.HandlerFunc
	GetMentorSchedule() echo.HandlerFunc
	Delete() echo.HandlerFunc
	CheckSchedule() echo.HandlerFunc
}

type ScheduleService interface {
	PostSchedule(newSchedule Core) error
	GetMentorSchedule(mentorID uint) ([]Core, error)
	Delete(mentorID, scheduleID uint) error
	CheckSchedule(input Core) error
}

type ScheduleData interface {
	PostSchedule(newClass Core) error
	GetMentorSchedule(mentorID uint) ([]Core, error)
	Delete(mentorID, scheduleID uint) error
	CheckSchedule(input Core) (int64, error)
	DetailSchedule(scheduleID uint) (Core, error)
}
