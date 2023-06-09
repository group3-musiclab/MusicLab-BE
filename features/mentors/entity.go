package mentors

import (
	"mime/multipart"
	"musiclab-be/features/instruments"
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
	TokenOauth           string
	Sex                  string
	Phone                string `validate:"number,max=12"`
	Address              string
	Instagram            string
	About                string
	AvgRating            float32
	CountReviews         int64
	MentorInstrument     instruments.MentorInstrumentCore
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

type MentorFilter struct {
	Name          string
	Instrument    int
	Genre         int
	Rating        float64
	Qualification string
}

type MentorsHandler interface {
	GetAll() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	GetProfileByIdParam() echo.HandlerFunc
	UpdateData() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
	AddCredential() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByRating() echo.HandlerFunc
}

type MentorService interface {
	SelectAllByRating() ([]Core, error)
	SelectAll(page, limit int, filter MentorFilter) ([]Core, error)
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	UpdatePassword(mentorID uint, input Core) error
	InsertCredential(input CredentialCore) error
	Delete(mentorID uint) error
}

type MentorData interface {
	SelectAllByRating() ([]Core, error)
	SelectAll(limit, offset int, filter MentorFilter) ([]Core, error)
	SelectProfile(mentorID uint) (Core, error)
	UpdateData(mentorID uint, input Core) error
	InsertCredential(input CredentialCore) error
	Delete(mentorID uint) error
}
