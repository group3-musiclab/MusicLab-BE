package handler

import (
	"log"
	"musiclab-be/features/schedules"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type scheduleControll struct {
	srv schedules.ScheduleService
}

// Delete implements schedules.ScheduleHandler
func (*scheduleControll) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetMentorSchedule implements schedules.ScheduleHandler
func (*scheduleControll) GetMentorSchedule() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv schedules.ScheduleService) schedules.ScheduleHandler {
	return &scheduleControll{
		srv: srv,
	}
}

func (sc *scheduleControll) PostSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := PostSchedule{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		mentorSchedule := addPostScheduleToCore(input)
		mentorSchedule.MentorID = mentorID
		err = sc.srv.PostSchedule(mentorSchedule)
		if err != nil {
			log.Println("error running add mentor schedule service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success post a schedule",
		})

	}
}
