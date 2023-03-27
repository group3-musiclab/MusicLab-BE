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

type CheckSchedule struct {
	ClassID    uint `json:"class_id" form:"class_id"`
	ScheduleID uint `json:"schedule_id" form:"schedule_id"`
}

func checkScheduleToCore(data CheckSchedule) schedules.Core {
	return schedules.Core{
		ClassID: data.ClassID,
		Transaction: schedules.Transaction{
			ScheduleID: data.ScheduleID,
		},
	}
}
