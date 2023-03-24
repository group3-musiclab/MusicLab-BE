package services

import (
	"errors"
	"log"
	"musiclab-be/features/genres"
	"strings"
)

type genreUseCase struct {
	qry genres.GenreData
}

func New(gd genres.GenreData) genres.GenreService {
	return &genreUseCase{
		qry: gd,
	}
}

func (guu *genreUseCase) AddMentorGenre(newGenre genres.Core) (genres.Core, error) {
	res, err := guu.qry.AddMentorGenre(newGenre)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return genres.Core{}, errors.New(msg)
	}
	return res, nil
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
