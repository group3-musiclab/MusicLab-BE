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

func (guu *genreUseCase) AddMentorGenre(newGenre genres.Core) error {
	err := guu.qry.AddMentorGenre(newGenre)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
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

func (guc *genreUseCase) GetMentorGenre(mentorID uint) ([]genres.MentorGenreCore, error) {
	res, err := guc.qry.GetMentorGenre(mentorID)
	if err != nil {
		log.Println("query error", err.Error())
		return []genres.MentorGenreCore{}, errors.New("query error, problem with server")
	}
	return res, nil
}

func (guc *genreUseCase) Delete(mentorID, genreID uint) error {
	err := guc.qry.Delete(mentorID, genreID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, problem with server")
	}
	return nil
}
