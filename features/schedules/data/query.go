package data

import (
	"errors"
	"log"
	"musiclab-be/features/schedules"

	"gorm.io/gorm"
)

type scheduleQuery struct {
	db *gorm.DB
}

// Delete implements schedules.ScheduleData
func (*scheduleQuery) Delete(mentorID uint, scheduleID uint) error {
	panic("unimplemented")
}

// GetMentorSchedule implements schedules.ScheduleData
func (*scheduleQuery) GetMentorSchedule(mentorID uint) ([]schedules.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) schedules.ScheduleData {
	return &scheduleQuery{
		db: db,
	}
}

// PostSchedule implements schedules.ScheduleData
func (sq *scheduleQuery) PostSchedule(newClass schedules.Core) error {
	cnv := CoreToData(newClass)
	err := sq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}
