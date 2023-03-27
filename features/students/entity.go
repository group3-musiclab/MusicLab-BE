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
	Name                 string `validate:"required,max=50"`
	Email                string `validate:"required,email,max=50"`
	Password             string `validate:"required,min=3"`
	NewPassword          string
	ConfirmationPassword string
	Role                 string
	Sex                  string
	Phone                string `validate:"number,max=12"`
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
	SelectProfile(studentID uint) (Core, error)
	UpdateData(studentID uint, input Core) error
	UpdatePassword(studentID uint, input Core) error
	Delete(studentID uint) error
}

type StudentData interface {
	SelectProfile(studentID uint) (Core, error)
	UpdateData(studentID uint, input Core) error
	Delete(studentID uint) error
}
