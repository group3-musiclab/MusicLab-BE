package reviews

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	MentorID  uint
	StudentID uint
	Rating    uint
	Comment   string
}

type ReviewHandler interface {
	PostMentorReview() echo.HandlerFunc
	GetMentorReview() echo.HandlerFunc
}

type ReviewService interface {
	PostMentorReview(newReview Core) error
	GetMentorReview(mentorID uint) ([]Core, error)
}

type ReviewData interface {
	PostMentorReview(newReview Core) error
	GetMentorReview(mentorID uint) ([]Core, error)
}
