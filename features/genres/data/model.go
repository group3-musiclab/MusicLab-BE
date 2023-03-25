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
	Genre    Genre
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

func GenreToCore(data Genre) genres.Core {
	return genres.Core{
		ID:   data.ID,
		Name: data.Name,
	}
}

func CoreToGenreData(data genres.Core) Genre {
	return Genre{
		Model: gorm.Model{ID: data.ID},
		Name:  data.Name,
	}
}

func MentorGenreCoreToModel(data genres.MentorGenreCore) MentorGenre {
	return MentorGenre{
		MentorID: data.MentorID,
		GenreID:  data.GenreID,
	}
}

func MentorGenreModelToCore(data MentorGenre) genres.MentorGenreCore {
	return genres.MentorGenreCore{
		ID:       data.ID,
		MentorID: data.MentorID,
		GenreID:  data.GenreID,
		Genre: genres.Core{
			Name: data.Genre.Name,
		},
	}
}

func MentorGenreListModelToCore(dataModel []MentorGenre) []genres.MentorGenreCore {
	var dataCore []genres.MentorGenreCore
	for _, v := range dataModel {
		dataCore = append(dataCore, MentorGenreModelToCore(v))
	}
	return dataCore
}
