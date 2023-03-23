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

func (gq *genreQuery) AddMentorGenre(mentorID uint, genreID uint, newGenre genres.Core) (genres.Core, error) {
	genre := Genre{}
	err := gq.db.Where("id=?", genreID).First(&genre).Error
	if err != nil {
		log.Println("query error", err.Error())
		return genres.Core{}, errors.New("server error")
	}
	cnv := CoreToData(newGenre)
	cnv.GenreID = genre.ID
	cnv.MentorID = mentorID
	err = gq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return genres.Core{}, errors.New("server error")
	}
	result := ToCore(cnv)
	return result, nil
}

// Delete implements genres.GenreData
func (gq *genreQuery) Delete(mentorID uint, genreID uint) error {
	panic("unimplemented")
}

// GetGenre implements genres.GenreData
func (gq *genreQuery) GetGenre() ([]genres.Core, error) {
	panic("unimplemented")
}

// GetMentorGenre implements genres.GenreData
func (gq *genreQuery) GetMentorGenre(mentorID uint) ([]genres.Core, error) {
	panic("unimplemented")
}
