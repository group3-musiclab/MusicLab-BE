package handler

import "musiclab-be/features/reviews"

type AddMentorReview struct {
	Rating  float64 `json:"rating" form:"rating"`
	Comment string  `json:"comment" form:"comment"`
}

func addMentorReviewToCore(data AddMentorReview) reviews.Core {
	return reviews.Core{
		Rating:  uint(data.Rating),
		Comment: data.Comment,
	}
}
