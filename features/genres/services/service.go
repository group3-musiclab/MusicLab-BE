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

// GetGenre implements genres.GenreService
func (guc *genreUseCase) GetGenre() ([]genres.Core, error) {
	res, err := guc.qry.GetGenre()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "genre not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []genres.Core{}, errors.New(msg)
	}
	return res, nil
}

func (guc *genreUseCase) GetMentorGenre(mentorID uint) ([]genres.Core, error) {
	res, err := guc.qry.GetMentorGenre(mentorID)
	if err != nil {
		log.Println("query error", err.Error())
		return []genres.Core{}, errors.New("query error, problem with server")
	}
	return res, nil
}

// Delete implements genres.GenreService
func (*genreUseCase) Delete(token interface{}, genreID uint) error {
	panic("unimplemented")
}
