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

// InsertCredential implements mentors.MentorService
func (muc *mentorUseCase) InsertCredential(input mentors.CredentialCore) error {
	errValidate := muc.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	url, errUpload := helper.GetUrlImagesFromAWS(input.CertificateFile)
	if errUpload != nil {
		return errors.New(consts.AWS_ErrorUpload)
	}

	input.Certificate = url

	errInsert := muc.qry.InsertCredential(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

// UpdatePassword implements mentors.MentorService
func (muc *mentorUseCase) UpdatePassword(mentorID uint, input mentors.Core) error {
	if input.Password == "" || input.NewPassword == "" || input.ConfirmationPassword == "" {
		return errors.New(consts.MENTOR_ErrorEmptyPassword)
	}

	dataCore, errSelect := muc.qry.SelectProfile(mentorID)
	if errSelect != nil {
		return errSelect
	}

	if !helper.CompareHashPassword(input.Password, dataCore.Password) {
		return errors.New(consts.AUTH_ErrorComparePassword)
	}

	if input.NewPassword != input.ConfirmationPassword {
		return errors.New(consts.AUTH_ErrorNewPassword)
	}

	hash, errHash := helper.HashPassword(input.NewPassword)
	if errHash != nil {
		return errors.New(consts.AUTH_ErrorHash)
	}

	input.Password = hash

	errUpdate := muc.qry.UpdateData(mentorID, input)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
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
