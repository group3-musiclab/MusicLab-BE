package data

import (
	"errors"
	"musiclab-be/features/chats"
	"musiclab-be/utils/consts"

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
func (cq *chatQuery) InsertChat(input chats.Core) error {
	dataModel := CoreToModel(input)
	txInsert := cq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return errors.New(consts.QUERY_ErrorInsertData)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

func New(db *gorm.DB) chats.ChatData {
	return &chatQuery{
		db: db,
	}
}
