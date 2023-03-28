package auth

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Avatar    string
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=3"`
	Role      string `validate:"required"`
	AvgRating float32
}

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GoogleLogin() echo.HandlerFunc
	GoogleCallback() echo.HandlerFunc
}

type AuthService interface {
	Register(newUser Core) error
	Login(user Core) (string, Core, error)
	CreateEvent(code string) error
}

type AuthData interface {
	RegisterMentor(newUser Core) error
	RegisterStudent(newUser Core) error
	LoginMentor(email string) (Core, error)
	LoginStudent(email string) (Core, error)
}
