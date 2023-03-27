package handler

import (
	"errors"
	"musiclab-be/features/schedules"
	"time"
)

type PostSchedule struct {
	Day       string `json:"day" form:"day"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}

func addPostScheduleToCore(data PostSchedule) schedules.Core {
	return schedules.Core{
		Day:       data.Day,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	}
}

type CheckSchedule struct {
	ClassID    uint   `json:"class_id" form:"class_id"`
	ScheduleID uint   `json:"schedule_id" form:"schedule_id"`
	StartDate  string `json:"start_date" form:"start_date"`
}

func checkScheduleToCore(data CheckSchedule) (schedules.Core, error) {
	StartDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return schedules.Core{}, errors.New("invalid start date format must YYYY-MM-DD")
	}
	return schedules.Core{
		ClassID: data.ClassID,
		Transaction: schedules.Transaction{
			ScheduleID: data.ScheduleID,
			StartDate:  StartDate,
		},
	}, nil
}
