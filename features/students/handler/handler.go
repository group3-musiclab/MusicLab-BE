package handler

import (
	"musiclab-be/features/students"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type studentControl struct {
	srv students.StudentService
}

// Delete implements students.StudentHandler
func (sc *studentControl) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetProfile implements students.StudentHandler
func (sc *studentControl) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		dataCore, err := sc.srv.SelectProfile(id)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.STUDENT_SuccessProfile, coreToProfileResponse(dataCore)))
	}
}

// UpdateData implements students.StudentHandler
func (sc *studentControl) UpdateData() echo.HandlerFunc {
	panic("unimplemented")
}

// UpdatePassword implements students.StudentHandler
func (sc *studentControl) UpdatePassword() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv students.StudentService) students.StudentHandler {
	return &studentControl{
		srv: srv,
	}
}
