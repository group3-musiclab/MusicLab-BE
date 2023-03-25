package data

import (
	"musiclab-be/features/reviews"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	MentorID  uint
	StudentID uint
	Rating    uint `gorm:"type:float"`
	Comment   string
}

func ToCore(data Review) reviews.Core {
	return reviews.Core{
		ID:        data.ID,
		MentorID:  data.MentorID,
		StudentID: data.StudentID,
		Rating:    data.Rating,
		Comment:   data.Comment,
	}
}

func CoreToData(data reviews.Core) Review {
	return Review{
		Model:     gorm.Model{ID: data.ID},
		MentorID:  data.MentorID,
		StudentID: data.StudentID,
		Rating:    data.Rating,
		Comment:   data.Comment,
	}
}
