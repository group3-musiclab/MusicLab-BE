package genres

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	MentorID uint
	GenreID  uint
}

type GenreHandler interface {
	AddMentorGenre() echo.HandlerFunc
	GetMentorGenre() echo.HandlerFunc
	GetGenre() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type GenreService interface {
	AddMentorGenre(token interface{}, genreID uint, newGenre Core) error
	GetMentorGenre(token interface{}) ([]Core, error)
	GetGenre() ([]Core, error)
	Delete(token interface{}, genreID uint) error
}

type GenreData interface {
	AddMentorGenre(mentorID uint, genreID uint, newGenre Core) error
	GetMentorGenre(mentorID uint) ([]Core, error)
	GetGenre() ([]Core, error)
	Delete(mentorID uint, genreID uint) error
}
