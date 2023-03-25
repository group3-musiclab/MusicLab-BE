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

// GetMentorClass implements classes.ClassData
func (*classQuery) GetMentorClass(mentorID uint) ([]classes.Core, error) {
	panic("unimplemented")
}

// GetMentorClassDetail implements classes.ClassData
func (*classQuery) GetMentorClassDetail(classID uint) (classes.Core, error) {
	panic("unimplemented")
}

// Update implements classes.ClassData
func (*classQuery) Update(mentorID uint, classID uint, updatedClass classes.Core) (classes.Core, error) {
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
