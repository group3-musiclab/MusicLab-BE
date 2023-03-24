package handler

import "musiclab-be/features/genres"

type ShowAllGenre struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ShowAllGenreResponse(data genres.Core) ShowAllGenre {
	return ShowAllGenre{
		ID:   data.ID,
		Name: data.Name,
	}
}
