package handler

import (
	"mime/multipart"
	"musiclab-be/features/classes"
)

type PostClass struct {
	Name        string `json:"name" form:"name"`
	Level       string `json:"level" form:"level"`
	Description string `json:"description" form:"description"`
	Syllabus    string `json:"syllabus" form:"syllabus"`
	ForWhom     string `json:"for_whom" form:"for_whom"`
	Requirement string `json:"requirement" form:"requirement"`
	Price       string `json:"price" form:"price"`
	Duration    uint   `json:"duration" form:"duration"`
	FileHeader  multipart.FileHeader
}

func addPostClassToCore(data PostClass) classes.Core {
	return classes.Core{
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
