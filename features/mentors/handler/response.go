package handler

import (
	"musiclab-be/features/mentors"
)

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
	CountReviews int64  `json:"count_reviews"`
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

type MentorResponse struct {
	ID        uint    `json:"id"`
	Avatar    string  `json:"avatar"`
	Name      string  `json:"name"`
	About     string  `json:"about"`
	Instagram string  `json:"instagram"`
	AvgRating float32 `json:"rating"`
}

func coreToResponse(data mentors.Core) MentorResponse {
	return MentorResponse{
		ID:        data.ID,
		Avatar:    data.Avatar,
		Name:      data.Name,
		About:     data.About,
		Instagram: data.Instagram,
		AvgRating: data.AvgRating,
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
