package services

import (
	"errors"
	"musiclab-be/features/instruments"

	"github.com/go-playground/validator/v10"
)

type instrumentUseCase struct {
	qry      instruments.InstrumentData
	validate *validator.Validate
}

// Delete implements instruments.InstrumentService
func (iuc *instrumentUseCase) Delete(mentorID uint, instrumentID uint) error {
	err := iuc.qry.Delete(mentorID, instrumentID)
	if err != nil {
		return err
	}
	return nil
}

// Insert implements instruments.InstrumentService
func (iuc *instrumentUseCase) Insert(input instruments.MentorInstrumentCore) error {
	errValidate := iuc.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validate: " + errValidate.Error())
	}

	errInsert := iuc.qry.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// SelectAll implements instruments.InstrumentService
func (iuc *instrumentUseCase) SelectAll() ([]instruments.Core, error) {
	data, err := iuc.qry.SelectAll()
	return data, err
}

// SelectAllByMentorID implements instruments.InstrumentService
func (iuc *instrumentUseCase) SelectAllByMentorID(mentorID uint) ([]instruments.MentorInstrumentCore, error) {
	panic("unimplemented")
}

func New(id instruments.InstrumentData) instruments.InstrumentService {
	return &instrumentUseCase{
		qry:      id,
		validate: validator.New(),
	}
}
