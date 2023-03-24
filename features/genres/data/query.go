package data

import (
	"errors"
	"log"
	"musiclab-be/features/genres"

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
func (gq *genreQuery) AddMentorGenre(newGenre genres.Core) (genres.Core, error) {
	cnv := CoreToData(newGenre)
	err := gq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return genres.Core{}, errors.New("server error")
	}
	result := ToCore(cnv)
	return result, nil
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

func (gq *genreQuery) Delete(mentorID uint, genreID uint) error {
	panic("unimplemented")
}

// GetMentorGenre implements genres.GenreData
func (gq *genreQuery) GetMentorGenre(mentorID uint) ([]genres.Core, error) {
	panic("unimplemented")
}
