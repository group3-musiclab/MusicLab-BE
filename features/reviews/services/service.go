package services

import (
	"errors"
	"log"
	"musiclab-be/features/reviews"
	"strings"
)

type reviewUseCase struct {
	qry reviews.ReviewData
}

func New(rd reviews.ReviewData) reviews.ReviewService {
	return &reviewUseCase{
		qry: rd,
	}
}

func (ruc *reviewUseCase) PostMentorReview(mentorID uint, newReview reviews.Core) error {
	err := ruc.qry.PostMentorReview(mentorID, newReview)
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

// GetMentorReview implements reviews.ReviewService
func (*reviewUseCase) GetMentorReview(mentorID uint) ([]reviews.Core, error) {
	panic("unimplemented")
}
