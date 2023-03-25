package handler

import (
	"mime/multipart"
	"musiclab-be/features/students"
)

type UpdateRequest struct {
	AvatarFile multipart.FileHeader `json:"avatar_file" form:"avatar_file"`
	Name       string               `json:"name" form:"name"`
	Email      string               `json:"email" form:"email"`
	Sex        string               `json:"sex" form:"sex"`
	Phone      string               `json:"phone" form:"phone"`
	Address    string               `json:"address" form:"address"`
}

func updateRequestToCore(data UpdateRequest) students.Core {
	return students.Core{
		AvatarFile: data.AvatarFile,
		Name:       data.Name,
		Email:      data.Email,
		Sex:        data.Sex,
		Phone:      data.Phone,
		Address:    data.Address,
	}
}

type UpdatePasswordRequest struct {
	Password             string `json:"old_password" form:"old_password"`
	NewPassword          string `json:"new_password" form:"new_password"`
	ConfirmationPassword string `json:"confirmation_password" form:"confirmation_password"`
}

func updatePasswordRequestToCore(data UpdatePasswordRequest) students.Core {
	return students.Core{
		Password:             data.Password,
		NewPassword:          data.NewPassword,
		ConfirmationPassword: data.ConfirmationPassword,
	}
}
