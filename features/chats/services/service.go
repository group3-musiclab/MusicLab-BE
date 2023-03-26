package services

import (
	"musiclab-be/features/chats"

	"github.com/go-playground/validator/v10"
)

type chatUseCase struct {
	qry      chats.ChatData
	validate *validator.Validate
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
func (cuc *chatUseCase) InsertChat(chats.Core) error {
	panic("unimplemented")
}

func New(cd chats.ChatData) chats.ChatService {
	return &chatUseCase{
		qry:      cd,
		validate: validator.New(),
	}
}
