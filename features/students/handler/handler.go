package handler

import (
	"musiclab-be/features/students"

	"github.com/labstack/echo/v4"
)

type studentControl struct {
	srv students.StudentService
}

// Delete implements students.StudentHandler
func (*studentControl) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetProfile implements students.StudentHandler
func (*studentControl) GetProfile() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdateData implements students.StudentHandler
func (*studentControl) UpdateData() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdatePassword implements students.StudentHandler
func (*studentControl) UpdatePassword() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv students.StudentService) students.StudentHandler {
	return &studentControl{
		srv: srv,
	}
}
