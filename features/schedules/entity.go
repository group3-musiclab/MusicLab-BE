package schedules

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	MentorID  uint
	Day       string
	StartTime time.Time
	EndTime   time.Time
}

type ScheduleHandler interface {
	PostSchedule() echo.HandlerFunc
	GetMentorSchedule() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ScheduleService interface {
	PostSchedule(newSchedule Core) error
	GetMentorSchedule(mentorID uint) ([]Core, error)
	Delete(mentorID, scheduleID uint) error
}

type ScheduleData interface {
	PostClass(newClass Core) error
	GetMentorSchedule(mentorID uint) ([]Core, error)
	Delete(mentorID, scheduleID uint) error
}
