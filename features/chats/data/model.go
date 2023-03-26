package data

import (
	"musiclab-be/features/chats"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	StudentID  uint
	MentorID   uint
	SenderName string
	Chat       string
}

func ModelToCore(data Chat) chats.Core {
	return chats.Core{
		ID:         data.ID,
		StudentID:  data.StudentID,
		MentorID:   data.MentorID,
		SenderName: data.SenderName,
		Chat:       data.Chat,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func CoreToModel(data chats.Core) Chat {
	return Chat{
		StudentID:  data.StudentID,
		MentorID:   data.MentorID,
		SenderName: data.SenderName,
		Chat:       data.Chat,
	}
}

func ListModelToCore(dataModel []Chat) []chats.Core {
	var dataCore []chats.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
