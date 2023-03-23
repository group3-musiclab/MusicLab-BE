package data

import "gorm.io/gorm"

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
