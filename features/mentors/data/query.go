package data

import (
	"musiclab-be/features/mentors"

	"gorm.io/gorm"
)

type mentorQuery struct {
	db *gorm.DB
}

// Login implements mentors.MentorData
func (*mentorQuery) Login(email string) (mentors.Core, error) {
	panic("unimplemented")
}

// Register implements mentors.MentorData
func (*mentorQuery) Register(newUser mentors.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) mentors.MentorData {
	return &mentorQuery{
		db: db,
	}
}
