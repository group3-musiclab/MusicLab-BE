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
func (*reviewQuery) GetMentorReview(mentorID uint) ([]reviews.Core, error) {
	panic("unimplemented")
}
