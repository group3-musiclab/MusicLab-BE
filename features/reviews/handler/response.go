package handler

type ShowAllMentorReview struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// func ShowAllMentorReviewResponse(data reviews.Core) ShowAllMentorReview {
// 	return ShowAllMentorReview{
// 		ID:   data.ID,
// 		Name: data.Name,
// 	}
// }
