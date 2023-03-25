package data

import (
	"errors"
	"log"
	"musiclab-be/features/reviews"

	"gorm.io/gorm"
)

type reviewQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reviews.ReviewData {
	return &reviewQuery{
		db: db,
	}
}

// PostMentorReview implements reviews.ReviewData
func (rq *reviewQuery) PostMentorReview(mentorID uint, newReview reviews.Core) error {
	res := Mentor{}
	err := rq.db.Where("id = ?", mentorID).First(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	cnv := CoreToData(newReview)
	err = rq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}

// GetMentorReview implements reviews.ReviewData
func (rq *reviewQuery) GetMentorReview(mentorID uint) ([]reviews.Core, error) {
	res := []Review{}
	err := rq.db.Where("mentor_id = ?", mentorID).First(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []reviews.Core{}, errors.New("server error")
	}

	result := []reviews.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}

	return result, nil
}
