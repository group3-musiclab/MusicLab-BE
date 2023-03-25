package services

import "musiclab-be/features/reviews"

type reviewUseCase struct {
	qry reviews.ReviewData
}

func New(rd reviews.ReviewData) reviews.ReviewService {
	return &reviewUseCase{
		qry: rd,
	}
}
