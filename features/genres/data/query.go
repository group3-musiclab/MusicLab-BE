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

func (gq *genreQuery) GetMentorGenre(mentorID uint) ([]genres.Core, error) {
	res := []Genre{}
	err := gq.db.Where("mentor_id = ?", mentorID).Find(&res).Error

	if err != nil {
		log.Println("query error", err.Error())
		return []genres.Core{}, errors.New("server error")
	}
	result := []genres.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, GenreToCore(res[i]))
		// cari data user berdasarkan cart user_id

		genre := MentorGenre{}
		err = gq.db.Where("id = ?", res[i].ID).First(&genre).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []genres.Core{}, errors.New("server error")
		}
		err = gq.db.Where("mentor_id = ?", genre.MentorID).First(&genre).Error
		if err != nil {
			log.Println("query error", err.Error())
			return []genres.Core{}, errors.New("server error")
		}
	}
	return result, nil
}

func (gq *genreQuery) Delete(mentorID uint, genreID uint) error {
	panic("unimplemented")
}
