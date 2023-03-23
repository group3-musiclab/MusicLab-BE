package data

import (
	"musiclab-be/features/genres"

	"gorm.io/gorm"
)

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

func ToCore(data MentorGenre) genres.Core {
	return genres.Core{
		ID:       data.ID,
		MentorID: data.MentorID,
		GenreID:  data.GenreID,
	}
}

func CoreToData(data genres.Core) MentorGenre {
	return MentorGenre{
		Model:    gorm.Model{ID: data.ID},
		MentorID: data.MentorID,
		GenreID:  data.GenreID,
	}
}
