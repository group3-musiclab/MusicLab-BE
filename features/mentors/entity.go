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
	AvgRating float32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MentorsHandler interface {
	GetProfile() echo.HandlerFunc
}

type MentorService interface {
	SelectProfile(newUser Core) error
}

type MentorData interface {
	SelectProfile(newUser Core) error
}
