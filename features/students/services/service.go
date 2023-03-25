package services

import (
	"musiclab-be/features/students"

	"github.com/go-playground/validator/v10"
)

type studentUseCase struct {
	qry      students.StudentData
	validate *validator.Validate
}

// Delete implements students.StudentService
func (*studentUseCase) Delete(mentorID uint) error {
	panic("unimplemented")
}

// SelectProfile implements students.StudentService
func (*studentUseCase) SelectProfile(mentorID uint) (students.Core, error) {
	panic("unimplemented")
}

// UpdateData implements students.StudentService
func (*studentUseCase) UpdateData(mentorID uint, input students.Core) error {
	panic("unimplemented")
}

// UpdatePassword implements students.StudentService
func (*studentUseCase) UpdatePassword(mentorID uint, input students.Core) error {
	panic("unimplemented")
}

func New(sd students.StudentData) students.StudentService {
	return &studentUseCase{
		qry:      sd,
		validate: validator.New(),
	}
}
