package data

import (
	"musiclab-be/features/students"

	"gorm.io/gorm"
)

type studentQuery struct {
	db *gorm.DB
}

// Delete implements students.StudentData
func (*studentQuery) Delete(mentorID uint) error {
	panic("unimplemented")
}

// SelectProfile implements students.StudentData
func (*studentQuery) SelectProfile(mentorID uint) (students.Core, error) {
	panic("unimplemented")
}

// UpdateData implements students.StudentData
func (*studentQuery) UpdateData(mentorID uint, input students.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) students.StudentData {
	return &studentQuery{
		db: db,
	}
}
