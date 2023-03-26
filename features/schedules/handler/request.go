package handler

import "musiclab-be/features/schedules"

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
