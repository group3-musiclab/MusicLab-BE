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

// GetByRating implements mentors.MentorsHandler
func (mc *mentorControl) GetByRating() echo.HandlerFunc {
	return func(c echo.Context) error {
		dataCore, err := mc.srv.SelectAllByRating()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.MENTOR_SuccessGetAll, listCoreToResponse(dataCore)))
	}
}

// GetAll implements mentors.MentorsHandler
func (mc *mentorControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var page int = 1
		pageParam := c.QueryParam("page")
		if pageParam != "" {
			pageConv, errConv := strconv.Atoi(pageParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				page = pageConv
			}
		}

		var limit int = 6
		limitParam := c.QueryParam("limit")
		if limitParam != "" {
			limitConv, errConv := strconv.Atoi(limitParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				limit = limitConv
			}
		}

		rating := float64(0)
		ratingParam := c.QueryParam("rating")
		if ratingParam != "" {
			ratingConv, errConv := strconv.Atoi(ratingParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidRatingParam))
			}
			rating = float64(ratingConv)
		}

		instrument := 0
		instrumentParam := c.QueryParam("instrument")
		if instrumentParam != "" {
			instrumentConv, errConv := strconv.Atoi(instrumentParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidInstrumentParam))
			}
			instrument = instrumentConv
		}

		genre := 0
		genreParam := c.QueryParam("genre")
		if genreParam != "" {
			genreConv, errConv := strconv.Atoi(genreParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidGenreParam))
			}
			genre = genreConv
		}

		mentorFilter := mentors.MentorFilter{
			Name:          c.QueryParam("name"),
			Instrument:    instrument,
			Genre:         genre,
			Rating:        rating,
			Qualification: c.QueryParam("qualification"),
		}

		dataCore, err := mc.srv.SelectAll(page, limit, mentorFilter)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.MENTOR_SuccessGetAll, listCoreToResponse(dataCore)))
	}
}

// Delete implements mentors.MentorsHandler
func (mc *mentorControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		err := mc.srv.Delete(mentorID)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusOK, helper.Response(consts.MENTOR_SuccessDelete))
	}
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
