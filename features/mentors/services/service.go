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
func (*mentorUseCase) SelectProfile(newUser mentors.Core) error {
	panic("unimplemented")
}

func New(md mentors.MentorData) mentors.MentorService {
	return &mentorUseCase{
		qry:      md,
		validate: validator.New(),
	}
}
