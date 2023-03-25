package handler

import (
	"musiclab-be/features/instruments"

	"github.com/labstack/echo/v4"
)

type instrumentControl struct {
	srv instruments.InstrumentService
}

// Add implements instruments.InstrumentHandler
func (*instrumentControl) Add() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements instruments.InstrumentHandler
func (*instrumentControl) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAll implements instruments.InstrumentHandler
func (*instrumentControl) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAllByMentorID implements instruments.InstrumentHandler
func (*instrumentControl) GetAllByMentorID() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv instruments.InstrumentService) instruments.InstrumentHandler {
	return &instrumentControl{
		srv: srv,
	}
}
