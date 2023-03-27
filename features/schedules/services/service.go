package services

import (
	"errors"
	"log"
	"musiclab-be/features/classes"
	"musiclab-be/features/schedules"
	"strings"

	"github.com/go-playground/validator/v10"
)

type scheduleUseCase struct {
	qry      schedules.ScheduleData
	qryClass classes.ClassData
	validate *validator.Validate
}

func New(sd schedules.ScheduleData, cd classes.ClassData) schedules.ScheduleService {
	return &scheduleUseCase{
		qry:      sd,
		qryClass: cd,
		validate: validator.New(),
	}
}

// CheckSchedule implements schedules.ScheduleService
func (suc *scheduleUseCase) CheckSchedule(input schedules.Core) error {
	errValidate := suc.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	// get mentor id by class id
	coreClass, errGetClass := suc.qryClass.GetMentorClassDetail(input.ClassID)
	if errGetClass != nil {
		return errGetClass
	}

	input.MentorID = coreClass.MentorID

	// calculate end date
	endDate := input.Transaction.StartDate.AddDate(0, int(coreClass.Duration), 0)
	input.Transaction.EndDate = endDate

	// check availability
	rows, errCheck := suc.qry.CheckSchedule(input)
	if errCheck != nil {
		return errCheck
	}

	if rows != 0 {
		return errors.New("schedule not available")
	}

	return nil
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

func (suc *scheduleUseCase) Delete(mentorID uint, scheduleID uint) error {
	err := suc.qry.Delete(mentorID, scheduleID)

	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("data not found")
	}

	return nil
}
