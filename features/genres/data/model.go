package data

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name         string
	MentorGenres []MentorGenre
}

type MentorGenre struct {
	gorm.Model
	MentorID uint
	GenreID  uint
}
