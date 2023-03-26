package data

import (
	"musiclab-be/features/reviews"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	MentorID   uint
	StudentID  uint
	Rating     uint `gorm:"type:float"`
	Comment    string
	Mentor     Mentor
	Student    Student
	ReviewDate string
	Avatar     string
	Name       string
}

type Mentor struct {
	gorm.Model
	Name  string
	Email string
	Phone string
}

type Student struct {
	gorm.Model
	Avatar string
	Name   string
	Email  string
	Phone  string
}

func ToCore(data Review) reviews.Core {
	return reviews.Core{
		ID:         data.ID,
		MentorID:   data.MentorID,
		StudentID:  data.StudentID,
		Rating:     data.Rating,
		Comment:    data.Comment,
		Avatar:     data.Avatar,
		Name:       data.Name,
		ReviewDate: data.ReviewDate,
		Mentor: reviews.Mentor{
			ID:    data.Mentor.ID,
			Name:  data.Mentor.Name,
			Email: data.Mentor.Email,
			Phone: data.Mentor.Phone,
		},
		Student: reviews.Student{
			ID:     data.Student.ID,
			Avatar: data.Student.Avatar,
			Name:   data.Student.Name,
			Email:  data.Student.Email,
			Phone:  data.Student.Phone,
		},
		CreatedAt: data.CreatedAt,
	}
}

func CoreToData(data reviews.Core) Review {
	return Review{
		Model:      gorm.Model{ID: data.ID},
		MentorID:   data.MentorID,
		StudentID:  data.StudentID,
		Rating:     data.Rating,
		Comment:    data.Comment,
		ReviewDate: data.ReviewDate,
		Avatar:     data.Avatar,
		Name:       data.Name,
	}
}
