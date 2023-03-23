package data

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	MentorID  uint
	StudentID uint
	Rating    uint `gorm:"type:float"`
	Comment   string
}
