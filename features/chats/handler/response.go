package handler

import (
	"musiclab-be/features/chats"
)

type AllChatResponse struct {
	ID         uint   `json:"id"`
	SenderName string `json:"sender_name"`
	Chat       string `json:"chat"`
}

func coreToAllChatResponse(data chats.Core) AllChatResponse {
	return AllChatResponse{
		ID:         data.ID,
		SenderName: data.SenderName,
		Chat:       data.Chat,
	}
}

func listCoreToAllChatResponse(dataCore []chats.Core) []AllChatResponse {
	var dataResponse []AllChatResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToAllChatResponse(v))
	}
	return dataResponse
}

type ChatByStudentResponse struct {
	Avatar      string `json:"avatar"`
	StudentID   uint   `json:"student_id"`
	StudentName string `json:"student_name"`
}

func coreToChatByStudentResponse(data chats.Core) ChatByStudentResponse {
	return ChatByStudentResponse{
		Avatar:      data.Student.Avatar,
		StudentID:   data.StudentID,
		StudentName: data.Student.Name,
	}
}

func listCoreToChatByStudentResponse(dataCore []chats.Core) []ChatByStudentResponse {
	var dataResponse []ChatByStudentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToChatByStudentResponse(v))
	}
	return dataResponse
}
