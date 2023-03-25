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
func (suc *studentUseCase) Delete(studentID uint) error {
	panic("unimplemented")
}

// SelectProfile implements students.StudentService
func (suc *studentUseCase) SelectProfile(studentID uint) (students.Core, error) {
	dataCore, err := suc.qry.SelectProfile(studentID)
	if err != nil {
		return students.Core{}, err
	}
	return dataCore, nil
}

// UpdateData implements students.StudentService
func (suc *studentUseCase) UpdateData(studentID uint, input students.Core) error {
	panic("unimplemented")
}

// UpdatePassword implements students.StudentService
func (suc *studentUseCase) UpdatePassword(studentID uint, input students.Core) error {
	panic("unimplemented")
}

func New(sd students.StudentData) students.StudentService {
	return &studentUseCase{
		qry:      sd,
		validate: validator.New(),
	}
}
