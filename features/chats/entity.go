package chats

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	StudentID  uint `validate:"required"`
	MentorID   uint `validate:"required"`
	SenderName string
	Chat       string `validate:"required,max=500"`
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ChatHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByStudent() echo.HandlerFunc
}

type ChatService interface {
	InsertChat(input Core) error
	GetAll(mentorID, studentID uint) ([]Core, error)
	GetByStudent(mentorID uint) ([]Core, error)
}

type ChatData interface {
	InsertChat(input Core) error
	GetAll(mentorID, studentID uint) ([]Core, error)
	GetByStudent(mentorID uint) ([]Core, error)
}
