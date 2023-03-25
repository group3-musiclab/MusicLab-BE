package students

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                   uint
	AvatarFile           multipart.FileHeader
	Avatar               string
	Name                 string `validate:"required"`
	Email                string `validate:"required,email"`
	Password             string `validate:"required,min=3"`
	NewPassword          string
	ConfirmationPassword string
	Role                 string
	Sex                  string
	Phone                string
	Address              string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type StudentHandler interface {
	GetProfile() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type StudentService interface {
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	UpdatePassword(mentorID uint, input Core) error
	Delete(mentorID uint) error
}

type StudentData interface {
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	Delete(mentorID uint) error
}
