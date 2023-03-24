package services

import (
	"musiclab-be/features/mentors"

	"github.com/go-playground/validator/v10"
)

type mentorUseCase struct {
	qry      mentors.MentorData
	validate *validator.Validate
}

// SelectProfile implements mentors.MentorService
func (muc *mentorUseCase) SelectProfile(idMentor uint) (mentors.Core, error) {
	dataCore, err := muc.qry.SelectProfile(idMentor)
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
