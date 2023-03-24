package handler

import (
	"musiclab-be/features/auth"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authControll struct {
	srv auth.AuthService
}

func New(srv auth.AuthService) auth.AuthHandler {
	return &authControll{
		srv: srv,
	}
}

// Register implements auth.AuthHandler
func (ac *authControll) Register() echo.HandlerFunc {
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

func (ac *authControll) Login() echo.HandlerFunc {
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
