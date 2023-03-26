package services

import (
	"errors"
	"log"
	"musiclab-be/features/schedules"
	"strings"
)

type scheduleUseCase struct {
	qry schedules.ScheduleData
}

// Delete implements schedules.ScheduleService
func (*scheduleUseCase) Delete(mentorID uint, scheduleID uint) error {
	panic("unimplemented")
}

func New(sd schedules.ScheduleData) schedules.ScheduleService {
	return &scheduleUseCase{
		qry: sd,
	}
}

func (suc *scheduleUseCase) PostSchedule(newSchedule schedules.Core) error {
	err := suc.qry.PostSchedule(newSchedule)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}

func (suc *scheduleUseCase) GetMentorSchedule(mentorID uint) ([]schedules.Core, error) {
	res, err := suc.qry.GetMentorSchedule(mentorID)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "schedule not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []schedules.Core{}, errors.New(msg)
	}

	return res, nil
}
