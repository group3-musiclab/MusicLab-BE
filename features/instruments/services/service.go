package services

import (
	"musiclab-be/features/instruments"

	"github.com/go-playground/validator/v10"
)

type instrumentUseCase struct {
	qry      instruments.InstrumentData
	validate *validator.Validate
}

// Delete implements instruments.InstrumentService
func (*instrumentUseCase) Delete(instrumentID uint) error {
	panic("unimplemented")
}

// Insert implements instruments.InstrumentService
func (*instrumentUseCase) Insert(input instruments.MentorInstrumentCore) error {
	panic("unimplemented")
}

// SelectAll implements instruments.InstrumentService
func (*instrumentUseCase) SelectAll() ([]instruments.Core, error) {
	panic("unimplemented")
}

// SelectAllByMentorID implements instruments.InstrumentService
func (*instrumentUseCase) SelectAllByMentorID(mentorID uint) ([]instruments.MentorInstrumentCore, error) {
	panic("unimplemented")
}

func New(id instruments.InstrumentData) instruments.InstrumentService {
	return &instrumentUseCase{
		qry:      id,
		validate: validator.New(),
	}
}
