package classes

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	MentorID    uint
	Name        string
	Image       string
	Level       string
	Description string
	Syllabus    string
	Requirement string
	ForWhom     string
	Price       string
	Duration    uint
}

type ClassHandler interface {
	PostClass() echo.HandlerFunc
	GetMentorClass() echo.HandlerFunc
	GetMentorClassDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ClassService interface {
	PostClass(newClass Core) error
	GetMentorClass(mentorID uint) ([]Core, error)
	GetMentorClassDetail(classID uint) (Core, error)
	Update(mentorID, classID uint, updatedClass Core) (Core, error)
	Delete(mentorID, classID uint) error
}

type ClassData interface {
	PostClass(newClass Core) error
	GetMentorClass(mentorID uint) ([]Core, error)
	GetMentorClassDetail(classID uint) (Core, error)
	Update(mentorID, classID uint, updatedClass Core) (Core, error)
	Delete(mentorID, classID uint) error
}
