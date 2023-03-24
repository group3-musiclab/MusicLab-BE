package services

import (
	"errors"
	"musiclab-be/features/mentors"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"

	"github.com/go-playground/validator/v10"
)

type mentorUseCase struct {
	qry      mentors.MentorData
	validate *validator.Validate
}

// UpdatePassword implements mentors.MentorService
func (*mentorUseCase) UpdatePassword(mentorID uint, input mentors.Core) error {
	panic("unimplemented")
}

// UpdateData implements mentors.MentorService
func (muc *mentorUseCase) UpdateData(mentorID uint, input mentors.Core) error {
	errValidate := muc.validate.StructExcept(input, "Password")
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	url, errUpload := helper.GetUrlImagesFromAWS(input.AvatarFile)
	if errUpload != nil {
		return errors.New(consts.AWS_ErrorUpload)
	}

	input.Avatar = url

	errUpdate := muc.qry.UpdateData(mentorID, input)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// SelectProfile implements mentors.MentorService
func (muc *mentorUseCase) SelectProfile(mentorID uint) (mentors.Core, error) {
	dataCore, err := muc.qry.SelectProfile(mentorID)
	if err != nil {
		return mentors.Core{}, err
	}
	return dataCore, nil
}

func New(md mentors.MentorData) mentors.MentorService {
	return &mentorUseCase{
		qry:      md,
		validate: validator.New(),
	}
}
