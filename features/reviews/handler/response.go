package handler

import "musiclab-be/features/reviews"

type ShowAllMentorReview struct {
	ID        uint    `json:"id"`
	Avatar    string  `json:"avatar"`
	Name      string  `json:"name"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
	CreatedAt string  `json:"created_at"`
}

func ShowAllMentorReviewResponse(data reviews.Core) ShowAllMentorReview {
	CreatedAt := data.CreatedAt.Format("January 2, 2006")
	return ShowAllMentorReview{
		ID:        data.ID,
		Avatar:    data.Student.Avatar,
		Name:      data.Student.Name,
		Rating:    float64(data.Rating),
		Comment:   data.Comment,
		CreatedAt: CreatedAt,
	}
}
