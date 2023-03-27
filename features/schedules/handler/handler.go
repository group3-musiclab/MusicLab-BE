package handler

import (
	"log"
	"musiclab-be/features/schedules"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type scheduleControll struct {
	srv schedules.ScheduleService
}

func New(srv schedules.ScheduleService) schedules.ScheduleHandler {
	return &scheduleControll{
		srv: srv,
	}
}

// CheckSchedule implements schedules.ScheduleHandler
func (sc *scheduleControll) CheckSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CheckSchedule{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		errSchedule := sc.srv.CheckSchedule(checkScheduleToCore(input))
		if errSchedule != nil {
			return errSchedule
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "schedule available",
		})
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

func (sc *scheduleControll) GetMentorSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("mentor_id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		res, err := sc.srv.GetMentorSchedule(uint(mentorID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowMentorSchedule{}
		for _, val := range res {
			result = append(result, ShowMentorScheduleResponse(val))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success show all mentor schedule",
		})
	}
}

func (sc *scheduleControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		id := c.Param("schedule_id")
		scheduleID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}

		err := sc.srv.Delete(mentorID, uint(scheduleID))

		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes delete schedule",
		})
	}
}
