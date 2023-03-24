package handler

import "musiclab-be/features/mentors"

type ProfileResponse struct {
	ID           uint   `json:"id"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Sex          string `json:"sex"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Instagram    string `json:"instagram"`
	About        string `json:"about"`
	CountReviews int    `json:"count_reviews"`
}

func coreToProfileResponse(data mentors.Core) ProfileResponse {
	return ProfileResponse{
		ID:           data.ID,
		Avatar:       data.Avatar,
		Name:         data.Name,
		Email:        data.Email,
		Sex:          data.Sex,
		Phone:        data.Phone,
		Address:      data.Address,
		Instagram:    data.Instagram,
		About:        data.About,
		CountReviews: data.CountReviews,
	}
}
