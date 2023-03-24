package handler

import (
	"musiclab-be/features/mentors"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type mentorControl struct {
	srv mentors.MentorService
}

// UpdateData implements mentors.MentorsHandler
func (*mentorControl) UpdateData() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.Response("ok"))
	}
}

// GetProfileByIdParam implements mentors.MentorsHandler
func (mc *mentorControl) GetProfileByIdParam() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idConv, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		dataCore, err := mc.srv.SelectProfile(uint(idConv))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.MENTOR_SuccessGetProfile, coreToProfileResponse(dataCore)))
	}
}

// GetProfile implements mentors.MentorsHandler
func (mc *mentorControl) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		dataCore, err := mc.srv.SelectProfile(id)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.MENTOR_SuccessGetProfile, coreToProfileResponse(dataCore)))
	}
}

func New(srv mentors.MentorService) mentors.MentorsHandler {
	return &mentorControl{
		srv: srv,
	}
}
