package handler

import (
	"errors"
	"musiclab-be/features/classes"
)

type ShowAllMentorClass struct {
	ID    uint    `json:"id"`
	Image string  `json:"image"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ShowAllMentorClassResponse(data classes.Core) ShowAllMentorClass {
	return ShowAllMentorClass{
		ID:    data.ID,
		Image: data.Image,
		Name:  data.Name,
		Price: data.Price,
	}
}

type ShowMentorClassDetail struct {
	ID          uint    `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Level       string  `json:"level"`
	Description string  `json:"description"`
	Syllabus    string  `json:"syllabus"`
	ForWhom     string  `json:"for_whom"`
	Requirement string  `json:"requirement"`
	Price       float64 `json:"price"`
	Duration    uint    `json:"duration"`
}

func ShowMentorClassDetailResponse(data classes.Core) ShowMentorClassDetail {
	return ShowMentorClassDetail{
		ID:          data.ID,
		Image:       data.Image,
		Name:        data.Name,
		Level:       data.Level,
		Description: data.Description,
		Syllabus:    data.Syllabus,
		ForWhom:     data.ForWhom,
		Requirement: data.Requirement,
		Price:       data.Price,
		Duration:    data.Duration,
	}
}

func ConvertClassUpdateResponse(input classes.Core) (interface{}, error) {
	ResponseFilter := classes.Core{}
	ResponseFilter = input
	result := make(map[string]interface{})
	if ResponseFilter.ID != 0 {
		result["id"] = ResponseFilter.ID
	}
	if ResponseFilter.Image != "" {
		result["image"] = ResponseFilter.Image
	}
	if ResponseFilter.Name != "" {
		result["name"] = ResponseFilter.Name
	}
	if ResponseFilter.Level != "" {
		result["level"] = ResponseFilter.Level
	}
	if ResponseFilter.Description != "" {
		result["description"] = ResponseFilter.Description
	}
	if ResponseFilter.Syllabus != "" {
		result["syllabus"] = ResponseFilter.Syllabus
	}
	if ResponseFilter.ForWhom != "" {
		result["for_whom"] = ResponseFilter.ForWhom
	}
	if ResponseFilter.Requirement != "" {
		result["requirement"] = ResponseFilter.Requirement
	}
	if ResponseFilter.Price != 0 {
		result["price"] = ResponseFilter.Price
	}
	if ResponseFilter.Duration != 0 {
		result["duration"] = ResponseFilter.Duration
	}

	if len(result) < 1 {
		return classes.Core{}, errors.New("no data was change")
	}
	return result, nil
}
