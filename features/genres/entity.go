package genres

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	MentorID uint   `json:"mentor_id"`
	GenreID  uint   `json:"genre_id"`
}

type MentorGenreCore struct {
	ID       uint
	MentorID uint `validate:"required"`
	GenreID  uint `validate:"required"`
	Genre    Core
}

type GenreHandler interface {
	AddMentorGenre() echo.HandlerFunc
	GetGenre() echo.HandlerFunc
	GetMentorGenre() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type GenreService interface {
	AddMentorGenre(newGenre Core) error
	GetGenre() ([]Core, error)
	GetMentorGenre(mentorID uint) ([]MentorGenreCore, error)
	Delete(mentorID, genreID uint) error
}

type GenreData interface {
	AddMentorGenre(newGenre Core) error
	GetGenre() ([]Core, error)
	GetMentorGenre(mentorID uint) ([]MentorGenreCore, error)
	Delete(mentorID, genreID uint) error
}
