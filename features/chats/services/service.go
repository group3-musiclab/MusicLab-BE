package services

import (
	"errors"
	"musiclab-be/features/chats"
	"musiclab-be/features/mentors"
	"musiclab-be/features/students"

	"github.com/go-playground/validator/v10"
)

type chatUseCase struct {
	qry        chats.ChatData
	qrymentor  mentors.MentorData
	qrystudent students.StudentData
	validate   *validator.Validate
}

// GetAll implements chats.ChatService
func (cuc *chatUseCase) GetAll(mentorID uint, studentID uint) ([]chats.Core, error) {
	panic("unimplemented")
}

// GetByStudent implements chats.ChatService
func (cuc *chatUseCase) GetByStudent(mentorID uint) ([]chats.Core, error) {
	panic("unimplemented")
}

// InsertChat implements chats.ChatService
func (cuc *chatUseCase) InsertChat(input chats.Core) error {
	errValidate := cuc.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	if input.Role == "Mentor" {
		mentorCore, errMentor := cuc.qrymentor.SelectProfile(input.MentorID)
		if errMentor != nil {
			return errMentor
		}
		input.SenderName = mentorCore.Name
	} else if input.Role == "Student" {
		studentCore, errStudent := cuc.qrystudent.SelectProfile(input.StudentID)
		if errStudent != nil {
			return errStudent
		}
		input.SenderName = studentCore.Name
	}

	errInsert := cuc.qry.InsertChat(input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func New(cd chats.ChatData, md mentors.MentorData, sd students.StudentData) chats.ChatService {
	return &chatUseCase{
		qry:       cd,
		qrymentor: md,
		validate:  validator.New(),
	}
}
