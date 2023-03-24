package handler

import (
	"musiclab-be/features/mentors"

	"github.com/labstack/echo/v4"
)

type mentorControl struct {
	srv mentors.MentorService
}

// GetProfile implements mentors.MentorsHandler
func (*mentorControl) GetProfile() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv mentors.MentorService) mentors.MentorsHandler {
	return &mentorControl{
		srv: srv,
	}
}
