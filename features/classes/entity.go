package classes

import (
	"mime/multipart"

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
	Price       float64
	Duration    uint
}

// ss
type ClassHandler interface {
	PostClass() echo.HandlerFunc
	GetMentorClass() echo.HandlerFunc
	GetMentorClassDetail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ClassService interface {
	PostClass(fileData multipart.FileHeader, newClass Core) error
	GetMentorClass(mentorID uint, page, limit int) ([]Core, error)
	GetMentorClassDetail(classID uint) (Core, error)
	Update(mentorID, classID uint, fileData multipart.FileHeader, updatedClass Core) error
	Delete(mentorID, classID uint) error
}

type ClassData interface {
	PostClass(newClass Core) error
	GetMentorClass(mentorID uint, limit, offset int) ([]Core, error)
	GetMentorClassDetail(classID uint) (Core, error)
	Update(mentorID, classID uint, updatedClass Core) error
	Delete(mentorID, classID uint) error
}
