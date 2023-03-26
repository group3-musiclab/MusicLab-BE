package handler

import "musiclab-be/features/schedules"

type ShowMentorSchedule struct {
	ID        uint   `json:"id"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func ShowMentorScheduleResponse(data schedules.Core) ShowMentorSchedule {
	return ShowMentorSchedule{
		ID:        data.ID,
		Day:       data.Day,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	}
}
