package schedules

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	MentorID  uint
	Day       string
	StartTime string
	EndTime   string
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
	PostSchedule(newClass Core) error
	GetMentorSchedule(mentorID uint) ([]Core, error)
	Delete(mentorID, scheduleID uint) error
}
