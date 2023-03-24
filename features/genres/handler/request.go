package handler

import "musiclab-be/features/genres"

type AddMentorGenre struct {
	GenreID uint `json:"genre_id" form:"genre_id"`
}

func addMentorGenreToCore(data AddMentorGenre) genres.Core {
	return genres.Core{
		GenreID: data.GenreID,
	}
}
