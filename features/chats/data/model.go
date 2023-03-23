package data

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	StudentID  uint
	MentorID   uint
	SenderName string
	Chat       string
}
