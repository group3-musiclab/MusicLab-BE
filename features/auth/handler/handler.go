package handler

import (
	"musiclab-be/features/auth"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	oauthStateString = "random"
)

type authControl struct {
	srv       auth.AuthService
	googleApi helper.GoogleAPI
}

// CreateEvent implements auth.AuthHandler
func (ac *authControl) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := CreateEventRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		inputCore := createEventToCore(input)
		inputCore.ID = mentorID

		err := ac.srv.CreateEvent(inputCore)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, helper.Response(consts.AUTH_SuccessCreateEvent))
	}
}

// GoogleCallback implements auth.AuthHandler
func (ac *authControl) GoogleCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		state := c.QueryParam("state")
		if state != oauthStateString {
			return c.HTML(http.StatusUnauthorized, "invalid oauth state")
		}

		code := c.QueryParam("code")

		err := ac.srv.LoginOauth(code)
		if err != nil {
			return c.HTML(http.StatusInternalServerError, err.Error())
		}

		return c.HTML(http.StatusOK, "success login with google")
	}
}

// GoogleLogin implements auth.AuthHandler
func (ac *authControl) GoogleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, ac.googleApi.GetUrlAuth(oauthStateString))
	}
}

// Register implements auth.AuthHandler
func (ac *authControl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		errRegister := ac.srv.Register(registerToCore(input))
		if errRegister != nil {
			return c.JSON(helper.ErrorResponse(errRegister))
		}
		return c.JSON(http.StatusCreated, helper.Response(consts.AUTH_SuccessCreate))
	}
}

func (ac *authControl) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		token, res, errLogin := ac.srv.Login(loginToCore(input))
		if errLogin != nil {
			return c.JSON(helper.ErrorResponse(errLogin))
		}
		dataResponse := map[string]any{
			"id":    res.ID,
			"name":  res.Name,
			"role":  res.Role,
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.AUTH_SuccessLogin, dataResponse))
	}
}

func New(srv auth.AuthService, ga helper.GoogleAPI) auth.AuthHandler {
	return &authControl{
		srv:       srv,
		googleApi: ga,
	}
}
