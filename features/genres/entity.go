package genres

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	MentorID uint
	GenreID  uint
}

//s

type GenreHandler interface {
	AddMentorGenre() echo.HandlerFunc
	GetGenre(e echo.Context) echo.HandlerFunc
	Delete(e echo.Context) echo.HandlerFunc
}

type GenreService interface {
	AddMentorGenre(genreID uint, newGenre Core) (Core, error)
	GetGenre() ([]Core, error)
	Delete(token interface{}, genreID uint) error
}

type GenreData interface {
	AddMentorGenre(genreID uint, newGenre Core) (Core, error)
	GetGenre() ([]Core, error)
	Delete(mentorID uint, genreID uint) error
}
