package auth

import "github.com/labstack/echo/v4"

type Core struct {
	ID            uint
	Avatar        string
	Name          string `validate:"required,max=50"`
	Email         string `validate:"required,email,max=50"`
	Password      string `validate:"required,min=3"`
	Role          string `validate:"required"`
	TokenOauth    string
	AvgRating     float32
	TransactionID uint `validate:"required"`
}
type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GoogleOauth() echo.HandlerFunc
	GoogleLogin() echo.HandlerFunc
	GoogleCallback() echo.HandlerFunc
	CreateEvent() echo.HandlerFunc
}

type AuthService interface {
	Register(newUser Core) error
	Login(user Core) (string, Core, error)
	LoginOauth(input Core) (string, Core, error)
	RedirectGoogleCallback(code string) error
	CreateEvent(input Core) error
}

type AuthData interface {
	RegisterMentor(newUser Core) error
	RegisterStudent(newUser Core) error
	LoginMentor(email string) (Core, error)
	LoginStudent(email string) (Core, error)
	FindAccount(email string) (Core, error)
}
