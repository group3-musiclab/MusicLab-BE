package handler

import (
	"mime/multipart"
	"musiclab-be/features/mentors"
)

type UpdateRequest struct {
	ID         uint
	AvatarFile multipart.File
	Name       string
	Email      string
	Sex        string
	Phone      string
	Address    string
	Instagram  string
	About      string
}

func updateRequestToCore(data UpdateRequest) mentors.Core {
	return mentors.Core{
		AvatarFile: data.AvatarFile,
		Name:       data.Name,
		Email:      data.Email,
		Sex:        data.Sex,
		Phone:      data.Phone,
		Address:    data.Address,
		Instagram:  data.Instagram,
		About:      data.About,
	}
}
