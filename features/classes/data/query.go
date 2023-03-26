package data

import (
	"errors"
	"log"
	"musiclab-be/features/classes"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Delete implements classes.ClassData
func (*classQuery) Delete(mentorID uint, classID uint) error {
	panic("unimplemented")
}

func New(db *gorm.DB) classes.ClassData {
	return &classQuery{
		db: db,
	}
}

// PostClass implements classes.ClassData
func (cq *classQuery) PostClass(newClass classes.Core) error {
	cnv := CoreToData(newClass)
	err := cq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}

func (cq *classQuery) GetMentorClass(mentorID uint) ([]classes.Core, error) {
	res := []Class{}

	err := cq.db.Where("mentor_id = ?", mentorID).Find(&res).Error

	if err != nil {
		log.Println("query error", err.Error())
		return []classes.Core{}, errors.New("server error")
	}

	result := []classes.Core{}

	for _, val := range res {
		result = append(result, ToCore(val))
	}

	return result, nil
}

func (cq *classQuery) GetMentorClassDetail(classID uint) (classes.Core, error) {
	res := Class{}

	err := cq.db.Where("id = ?", classID).First(&res).Error

	if err != nil {
		log.Println("data not found", err.Error())
		return classes.Core{}, errors.New("data not found")
	}

	result := ToCore(res)

	return result, nil
}

func (cq *classQuery) Update(mentorID uint, classID uint, updatedClass classes.Core) (classes.Core, error) {
	cnv := CoreToData(updatedClass)
	cnv.ID = uint(classID)

	qry := cq.db.Where("id = ?", classID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return classes.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update class query error", err.Error())
		return classes.Core{}, errors.New("user not found")
	}
	return updatedClass, nil
}
