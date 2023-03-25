package data

import (
	"errors"
	"musiclab-be/features/students"
	"musiclab-be/utils/consts"

	"gorm.io/gorm"
)

type studentQuery struct {
	db *gorm.DB
}

// Delete implements students.StudentData
func (sq *studentQuery) Delete(studentID uint) error {
	tx := sq.db.Delete(&Student{}, studentID)
	if tx.Error != nil {
		return errors.New(consts.QUERY_NotFound)
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// SelectProfile implements students.StudentData
func (sq *studentQuery) SelectProfile(studentID uint) (students.Core, error) {
	dataModel := Student{}
	txSelect := sq.db.First(&dataModel, studentID)
	if txSelect.Error != nil {
		return students.Core{}, errors.New(consts.QUERY_NotFound)
	}

	return ModelToCore(dataModel), nil
}

// UpdateData implements students.StudentData
func (sq *studentQuery) UpdateData(studentID uint, input students.Core) error {
	dataModel := CoreToModel(input)
	tx := sq.db.Where("id = ?", studentID).Updates(&dataModel)
	if tx.Error != nil {
		return errors.New(consts.QUERY_ErrorUpdateData)
	}
	if tx.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

func New(db *gorm.DB) students.StudentData {
	return &studentQuery{
		db: db,
	}
}
