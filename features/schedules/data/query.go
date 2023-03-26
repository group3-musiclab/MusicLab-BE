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

func (sq *scheduleQuery) Delete(mentorID uint, scheduleID uint) error {
	getID := Schedule{}
	err := sq.db.Where("id = ? and mentor_id = ?", scheduleID, mentorID).First(&getID).Error
	if err != nil {
		log.Println("get class error : ", err.Error())
		return errors.New("failed to get class data")
	}

	qryDelete := sq.db.Delete(&Schedule{}, scheduleID)
	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("no rows affected")
		return errors.New("failed to delete book, data not found")
	}

	return nil
}
