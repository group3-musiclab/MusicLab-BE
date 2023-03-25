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
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		err := sc.srv.Delete(id)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.STUDENT_SuccessDelete))
	}
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
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		input := UpdateRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		checkFile, _, _ := c.Request().FormFile("avatar_file")
		if checkFile != nil {
			formHeader, err := c.FormFile("avatar_file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.Response(consts.HANDLER_ErrorFormFile))
			}
			input.AvatarFile = *formHeader
		}

		err := sc.srv.UpdateData(id, updateRequestToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.STUDENT_SuccessUpdateProfile))
	}
}

// UpdatePassword implements students.StudentHandler
func (sc *studentControl) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		input := UpdatePasswordRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		err := sc.srv.UpdatePassword(id, updatePasswordRequestToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.STUDENT_SuccessUpdatePassword))
	}
}

func New(srv students.StudentService) students.StudentHandler {
	return &studentControl{
		srv: srv,
	}
}
