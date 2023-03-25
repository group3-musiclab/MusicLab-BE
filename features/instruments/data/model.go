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
