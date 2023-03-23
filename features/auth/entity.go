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
	Login(user Core) (string, Core, error)
}

type AuthData interface {
	RegisterMentor(newUser Core) error
	RegisterStudent(newUser Core) error
	LoginMentor(email string) (Core, error)
	LoginStudent(email string) (Core, error)
}
