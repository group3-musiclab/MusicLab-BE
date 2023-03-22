package handler

import (
	"musiclab-be/features/auth"
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
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		err = ac.srv.Register(*ReqToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "already") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "is not min") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "validate: password length minimum 3 character"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{"message": "success create account"})

	}
}

// Login implements auth.AuthHandler
func (*authControll) Login() echo.HandlerFunc {
	return nil
}
