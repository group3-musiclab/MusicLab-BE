package mentors

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Avatar    string
	Name      string
	Email     string
	Password  string
	Role      string
	Sex       string
	Phone     string
	Address   string
	Instagram string
	Price     float64
	AvgRating float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MentorsHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type MentorService interface {
	Register(newUser Core) error
	Login(email, password string) (string, Core, error)
}

type MentorData interface {
	Register(newUser Core) error
	Login(email string) (Core, error)
}
