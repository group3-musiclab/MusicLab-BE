package handler

import (
	"musiclab-be/features/mentors"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mentorControl struct {
	srv mentors.MentorService
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
