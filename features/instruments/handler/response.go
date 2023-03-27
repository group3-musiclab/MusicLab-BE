package handler

import "musiclab-be/features/instruments"

type InstrumentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func coreToResponse(data instruments.Core) InstrumentResponse {
	return InstrumentResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}

func listCoreToResponse(dataCore []instruments.Core) []InstrumentResponse {
	var dataResponse []InstrumentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}

type MentorInstrumentResponse struct {
	InstrumentID uint   `json:"id"`
	Name         string `json:"name"`
}

func coreToMentorInstrumentResponse(data instruments.MentorInstrumentCore) MentorInstrumentResponse {
	return MentorInstrumentResponse{
		InstrumentID: data.InstrumentID,
		Name:         data.Instrument.Name,
	}
}

func listCoreToMentorInstrumentResponse(dataCore []instruments.MentorInstrumentCore) []MentorInstrumentResponse {
	var dataResponse []MentorInstrumentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToMentorInstrumentResponse(v))
	}
	return dataResponse
}
