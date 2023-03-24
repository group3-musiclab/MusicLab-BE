package handler

import (
	"musiclab-be/features/auth"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strings"

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
		return c.JSON(http.StatusCreated, consts.AUTH_SuccessCreate)
	}
}

func (ac *authControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		if input.Email == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email not allowed empty"})
		} else if input.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not allowed empty"})
		}

		token, res, err := ac.srv.Login(loginToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "password not match"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "account not registered"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToResponse(res),
			"token":   token,
			"message": "success login",
		})
	}
}
