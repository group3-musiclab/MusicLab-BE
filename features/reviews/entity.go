package reviews

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID         uint
	MentorID   uint
	StudentID  uint
	Rating     uint
	Comment    string
	Mentor     Mentor
	Student    Student
	ReviewDate string
	Avatar     string
	Name       string
	CreatedAt  time.Time
}

type Mentor struct {
	ID    uint
	Name  string
	Email string
	Phone string
}

type Student struct {
	ID     uint
	Avatar string
	Name   string
	Email  string
	Phone  string
}

type ReviewHandler interface {
	PostMentorReview() echo.HandlerFunc
	GetMentorReview() echo.HandlerFunc
}

type ReviewService interface {
	PostMentorReview(mentorID uint, newReview Core) error
	GetMentorReview(page, limit int, mentorID uint) ([]Core, error)
}

type ReviewData interface {
	PostMentorReview(mentorID uint, newReview Core) error
	GetMentorReview(limit, offset int, mentorID uint) ([]Core, error)
}
