package handler

import "musiclab-be/features/students"

type ProfileResponse struct {
	ID      uint   `json:"id"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Sex     string `json:"sex"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func coreToProfileResponse(data students.Core) ProfileResponse {
	return ProfileResponse{
		ID:      data.ID,
		Avatar:  data.Avatar,
		Name:    data.Name,
		Email:   data.Email,
		Sex:     data.Sex,
		Phone:   data.Phone,
		Address: data.Address,
	}
}
