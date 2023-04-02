package handler

import (
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
