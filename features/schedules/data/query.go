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

func (sq *scheduleQuery) GetMentorSchedule(mentorID uint) ([]schedules.Core, error) {
	res := []Schedule{}
	err := sq.db.Where("mentor_id = ?", mentorID).Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []schedules.Core{}, errors.New("server error")
	}
	result := []schedules.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}

	return result, nil
}
