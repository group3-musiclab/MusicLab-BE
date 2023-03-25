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

// AddCredential implements mentors.MentorsHandler
func (mc *mentorControl) AddCredential() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := CredentialRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		formHeader, err := c.FormFile("certificate_file")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(consts.HANDLER_ErrorFormFile))
		}
		input.CertificateFile = *formHeader

		dataCore := credentialRequestToCore(input)
		dataCore.MentorID = mentorID

		errInsert := mc.srv.InsertCredential(dataCore)
		if errInsert != nil {
			return c.JSON(helper.ErrorResponse(errInsert))
		}

		return c.JSON(http.StatusCreated, helper.Response(consts.MENTOR_SuccessAddCredential))

	}
}

// UpdatePassword implements mentors.MentorsHandler
func (mc *mentorControl) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := UpdatePasswordRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		err := mc.srv.UpdatePassword(mentorID, updatePasswordRequestToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.MENTOR_SuccessUpdatePassword))
	}
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

		checkFile, _, _ := c.Request().FormFile("avatar_file")
		if checkFile != nil {
			formHeader, err := c.FormFile("avatar_file")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.Response(consts.HANDLER_ErrorFormFile))
			}
			input.AvatarFile = *formHeader
		}

		err := mc.srv.UpdateData(mentorID, updateRequestToCore(input))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.MENTOR_SuccessUpdateProfile))
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
