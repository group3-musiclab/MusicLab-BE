package handler

import (
	"musiclab-be/features/auth"
)

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func registerToCore(data RegisterRequest) auth.Core {
	return auth.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func loginToCore(data LoginRequest) auth.Core {
	return auth.Core{
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

type CreateEventRequest struct {
	TransactionID uint `json:"transaction_id"`
}

func createEventToCore(data CreateEventRequest) auth.Core {
	return auth.Core{
		TransactionID: data.TransactionID,
	}
}
