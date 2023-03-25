package data

import (
	"errors"
	"log"
	"musiclab-be/features/genres"
	"musiclab-be/utils/consts"

	"gorm.io/gorm"
)

type genreQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) genres.GenreData {
	return &genreQuery{
		db: db,
	}
}

// s
func (gq *genreQuery) AddMentorGenre(newGenre genres.Core) error {
	cnv := CoreToData(newGenre)
	err := gq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}

func (gq *genreQuery) GetGenre() ([]genres.Core, error) {
	genre := []Genre{}
	err := gq.db.Find(&genre).Error
	if err != nil {
		log.Println("data not found", err.Error())
		return []genres.Core{}, errors.New("data not found")
	}

	result := []genres.Core{}
	for _, val := range genre {
		result = append(result, GenreToCore(val))
	}

	return result, nil
}

func (gq *genreQuery) GetMentorGenre(mentorID uint) ([]genres.MentorGenreCore, error) {
	res := []MentorGenre{}
	err := gq.db.Preload("Genre").Where("mentor_id = ?", mentorID).Find(&res)
	if err.Error != nil {
		return nil, errors.New("server error")
	}

	result := MentorGenreListModelToCore(res)
	return result, nil
}

func (gq *genreQuery) Delete(mentorID, genreID uint) error {
	data := MentorGenre{}

	err := gq.db.Where("mentor_id = ?", mentorID).Where("genre_id = ?", genreID).Delete(&data)
	if err.Error != nil {
		return errors.New(consts.QUERY_NotFound)
	}
	if err.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}
