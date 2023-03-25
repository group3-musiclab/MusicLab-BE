package mentors

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
	Instagram            string
	About                string
	AvgRating            float32
	CountReviews         int64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CredentialCore struct {
	ID              uint
	CertificateFile multipart.FileHeader `validate:"required"`
	MentorID        uint
	Name            string `validate:"required"`
	Type            string `validate:"required"`
	Certificate     string
}

type MentorsHandler interface {
	GetProfile() echo.HandlerFunc
	GetProfileByIdParam() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	AddCredential() echo.HandlerFunc
}

type MentorService interface {
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	UpdatePassword(mentorID uint, input Core) error
	InsertCredential(input CredentialCore) error
}

type MentorData interface {
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	InsertCredential(input CredentialCore) error
}
