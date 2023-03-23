package data

import (
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
	Price        string `gorm:"type:float not null"`
	Duration     uint
	Transactions []_modelTransaction.Transaction
}
