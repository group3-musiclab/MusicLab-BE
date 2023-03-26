package data

import (
	"musiclab-be/features/chats"

	"gorm.io/gorm"
)

type chatQuery struct {
	db *gorm.DB
}

// GetAll implements chats.ChatData
func (cq *chatQuery) GetAll(mentorID uint, studentID uint) ([]chats.Core, error) {
	panic("unimplemented")
}

// GetByStudent implements chats.ChatData
func (cq *chatQuery) GetByStudent(mentorID uint) ([]chats.Core, error) {
	panic("unimplemented")
}

// InsertChat implements chats.ChatData
func (cq *chatQuery) InsertChat(chats.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) chats.ChatData {
	return &chatQuery{
		db: db,
	}
}
