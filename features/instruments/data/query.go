package data

import (
	"errors"
	"musiclab-be/features/instruments"
	"musiclab-be/utils/consts"

	"gorm.io/gorm"
)

type instrumentQuery struct {
	db *gorm.DB
}

// Delete implements instruments.InstrumentData
func (iq *instrumentQuery) Delete(instrumentID uint) error {
	panic("unimplemented")
}

// Insert implements instruments.InstrumentData
func (iq *instrumentQuery) Insert(input instruments.MentorInstrumentCore) error {
	dataModel := MentorInstrumentCoreToModel(input)
	txInsert := iq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return errors.New(consts.QUERY_ErrorInsertData)
	}
	if txInsert.RowsAffected == 0 {
		return errors.New(consts.QUERY_NoRowsAffected)
	}
	return nil
}

// SelectAll implements instruments.InstrumentData
func (iq *instrumentQuery) SelectAll() ([]instruments.Core, error) {
	var dataModel []Instrument
	tx := iq.db.Find(&dataModel)
	if tx.Error != nil {
		return nil, errors.New(consts.QUERY_ErrorReadData)
	}
	classCoreAll := ListModelToCore(dataModel)
	return classCoreAll, nil
}

// SelectAllByMentorID implements instruments.InstrumentData
func (iq *instrumentQuery) SelectAllByMentorID(mentorID uint) ([]instruments.MentorInstrumentCore, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) instruments.InstrumentData {
	return &instrumentQuery{
		db: db,
	}
}
