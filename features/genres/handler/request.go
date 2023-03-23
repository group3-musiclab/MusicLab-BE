package handler

import "musiclab-be/features/genres"

type AddMentorGenre struct {
	GenreID uint `json:"genre_id" form:"genre_id"`
}

func ReqToCore(data interface{}) *genres.Core {
	res := genres.Core{}

	switch data.(type) {
	case AddMentorGenre:
		cnv := data.(AddMentorGenre)
		res.GenreID = cnv.GenreID
	default:
		return nil
	}
	return &res
}
