package services

import (
	"musiclab-be/features/genres"
)

type genreUseCase struct {
	qry genres.GenreData
}

func New(gd genres.GenreData) genres.GenreService {
	return &genreUseCase{
		qry: gd,
	}
}

func (guu *genreUseCase) AddMentorGenre(genreID uint, newGenre genres.Core) (genres.Core, error) {

}

// Delete implements genres.GenreService
func (*genreUseCase) Delete(token interface{}, genreID uint) error {
	panic("unimplemented")
}

// GetGenre implements genres.GenreService
func (*genreUseCase) GetGenre() ([]genres.Core, error) {
	panic("unimplemented")
}

// GetMentorGenre implements genres.GenreService
func (*genreUseCase) GetMentorGenre(token interface{}) ([]genres.Core, error) {
	panic("unimplemented")
}
