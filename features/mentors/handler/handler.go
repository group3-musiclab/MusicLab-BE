package handler

import (
	"mime/multipart"
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
func (mc *mentorControl) UpdateData() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := UpdateRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		var blobFile multipart.File
		checkFile, _, _ := c.Request().FormFile("avatar_file")
		if checkFile != nil {
			file, errFile := c.FormFile("avatar_file")
			if errFile != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorFormFile))
			}

			var errBlob error
			blobFile, errBlob = file.Open()
			if errBlob != nil {
				return c.JSON(http.StatusNotFound, helper.Response(consts.HANDLER_ErrorBlobFile))
			}
		}

		dataCore := updateRequestToCore(input)
		dataCore.AvatarFile = blobFile

		err := mc.srv.UpdateData(mentorID, dataCore)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed add user")
		}
		return c.JSON(http.StatusCreated, "success add user")
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
