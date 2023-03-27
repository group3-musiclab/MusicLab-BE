package data

import (
	"musiclab-be/features/classes"
	_modelTransaction "musiclab-be/features/transactions/data"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	MentorID     uint
	Name         string
	Image        string
	Level        string `gorm:"type:enum('Basic','Intermediate','Advanced')"`
	Description  string
	Syllabus     string
	Requirement  string
	ForWhom      string
	Price        float64 `gorm:"type:float not null"`
	Duration     uint
	Transactions []_modelTransaction.Transaction
}

func ToCore(data Class) classes.Core {
	return classes.Core{
		ID:          data.ID,
		MentorID:    data.MentorID,
		Name:        data.Name,
		Image:       data.Image,
		Level:       data.Level,
		Description: data.Description,
		Syllabus:    data.Syllabus,
		Requirement: data.Requirement,
		ForWhom:     data.ForWhom,
		Price:       data.Price,
		Duration:    data.Duration,
	}
}

func CoreToData(data classes.Core) Class {
	return Class{
		Model:       gorm.Model{ID: data.ID},
		MentorID:    data.MentorID,
		Name:        data.Name,
		Image:       data.Image,
		Level:       data.Level,
		Description: data.Description,
		Syllabus:    data.Syllabus,
		Requirement: data.Requirement,
		ForWhom:     data.ForWhom,
		Price:       data.Price,
		Duration:    data.Duration,
	}
}
