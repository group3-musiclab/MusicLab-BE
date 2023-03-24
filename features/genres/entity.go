package genres

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	MentorID uint   `json:"mentor_id"`
	GenreID  uint   `json:"genre_id"`
}

//s

type GenreHandler interface {
	AddMentorGenre() echo.HandlerFunc
	GetGenre() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type GenreService interface {
	AddMentorGenre(newGenre Core) (Core, error)
	GetGenre() ([]Core, error)
	Delete(token interface{}, genreID uint) error
}

type GenreData interface {
	AddMentorGenre(newGenre Core) (Core, error)
	GetGenre() ([]Core, error)
	Delete(mentorID uint, genreID uint) error
}
