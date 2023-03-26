package handler

import (
	"musiclab-be/features/chats"
	"musiclab-be/features/mentors"
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

func listCoreToResponse(dataCore []mentors.Core) []MentorResponse {
	var dataResponse []MentorResponse
	for _, v := range dataCore {
		if len(v.About) > 79 {
			v.About = v.About[:79] + "..."
		}
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
