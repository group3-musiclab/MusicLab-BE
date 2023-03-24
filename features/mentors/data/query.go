package data

import (
	"musiclab-be/features/mentors"

	"gorm.io/gorm"
)

type mentorQuery struct {
	db *gorm.DB
}

// SelectProfile implements mentors.MentorData
func (*mentorQuery) SelectProfile(idMentor uint) (mentors.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) mentors.MentorData {
	return &mentorQuery{
		db: db,
	}
}
