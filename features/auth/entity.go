package auth

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Avatar   string
	Name     string
	Email    string
	Password string
	Role     string
}

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type AuthService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
}

type AuthData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
}
