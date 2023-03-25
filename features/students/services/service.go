package services

import (
	"errors"
	"musiclab-be/features/students"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"

	"github.com/go-playground/validator/v10"
)

type studentUseCase struct {
	qry      students.StudentData
	validate *validator.Validate
}

// Delete implements students.StudentService
func (suc *studentUseCase) Delete(studentID uint) error {
	err := suc.qry.Delete(studentID)
	if err != nil {
		return err
	}
	return nil
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
	errValidate := suc.validate.StructExcept(input, "Password")
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	url, errUpload := helper.GetUrlImagesFromAWS(input.AvatarFile)
	if errUpload != nil {
		return errors.New(consts.AWS_ErrorUpload)
	}

	input.Avatar = url

	errUpdate := suc.qry.UpdateData(studentID, input)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
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
