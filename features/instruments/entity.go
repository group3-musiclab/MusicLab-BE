package instruments

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MentorInstrumentCore struct {
	ID           uint
	MentorID     uint `validate:"required"`
	InstrumentID uint `validate:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type InstrumentHandler interface {
	GetAll() echo.HandlerFunc
	GetAllByMentorID() echo.HandlerFunc
	Add() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type InstrumentService interface {
	SelectAll() ([]Core, error)
	SelectAllByMentorID(mentorID uint) ([]MentorInstrumentCore, error)
	Insert(input MentorInstrumentCore) error
	Delete(instrumentID uint) error
}

type InstrumentData interface {
	SelectAll() ([]Core, error)
	SelectAllByMentorID(mentorID uint) ([]MentorInstrumentCore, error)
	Insert(input MentorInstrumentCore) error
	Delete(instrumentID uint) error
}
