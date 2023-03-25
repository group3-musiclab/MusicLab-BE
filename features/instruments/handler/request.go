package handler

import (
	"musiclab-be/features/instruments"
)

type MentorInstrumentRequest struct {
	InstrumentID uint `json:"instrument_id" form:"instrument_id"`
}

func insertRequestToCore(data MentorInstrumentRequest) instruments.MentorInstrumentCore {
	return instruments.MentorInstrumentCore{
		InstrumentID: data.InstrumentID,
	}
}
