package data

import (
	"musiclab-be/features/instruments"

	"gorm.io/gorm"
)

type Instrument struct {
	gorm.Model
	Name              string
	MentorInstruments []MentorInstrument
}

type MentorInstrument struct {
	gorm.Model
	MentorID     uint
	InstrumentID uint
}

func CoreToModel(data instruments.Core) Instrument {
	return Instrument{
		Name: data.Name,
	}
}

func ModelToCore(data Instrument) instruments.Core {
	return instruments.Core{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ListModelToCore(dataModel []Instrument) []instruments.Core {
	var dataCore []instruments.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}

func MentorInstrumentCoreToModel(data instruments.MentorInstrumentCore) MentorInstrument {
	return MentorInstrument{
		MentorID:     data.MentorID,
		InstrumentID: data.InstrumentID,
	}
}

func MentorInstrumentModelToCore(data MentorInstrument) instruments.MentorInstrumentCore {
	return instruments.MentorInstrumentCore{
		ID:           data.ID,
		MentorID:     data.MentorID,
		InstrumentID: data.InstrumentID,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func MentorInstrumentListModelToCore(dataModel []MentorInstrument) []instruments.MentorInstrumentCore {
	var dataCore []instruments.MentorInstrumentCore
	for _, v := range dataModel {
		dataCore = append(dataCore, MentorInstrumentModelToCore(v))
	}
	return dataCore
}
