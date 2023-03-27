package handler

import "musiclab-be/features/chats"

type AddChatRequest struct {
	MentorID  uint   `json:"mentor_id" form:"mentor_id"`
	StudentID uint   `json:"student_id" form:"student_id"`
	Chat      string `json:"chat" form:"chat"`
}

func addRequestToCore(data AddChatRequest) chats.Core {
	return chats.Core{
		StudentID: data.StudentID,
		MentorID:  data.MentorID,
		Chat:      data.Chat,
	}
}
