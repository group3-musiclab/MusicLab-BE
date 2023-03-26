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
	var dataModel []Chat
	txSelect := cq.db.Where("mentor_id", mentorID).Where("student_id", studentID).Find(&dataModel)
	if txSelect.Error != nil {
		return nil, errors.New(consts.QUERY_ErrorReadData)
	}
	return ListModelToCore(dataModel), nil
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
