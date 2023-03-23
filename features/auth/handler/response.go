package handler

import "musiclab-be/features/auth"

type AuthResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

func ToResponse(data auth.Core) AuthResponse {
	return AuthResponse{
		ID:   data.ID,
		Name: data.Name,
		Role: data.Role,
	}
}
