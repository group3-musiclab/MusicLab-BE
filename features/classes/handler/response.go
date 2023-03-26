package handler

import "musiclab-be/features/classes"

type ShowAllMentorClass struct {
	ID    uint   `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func ShowAllMentorClassResponse(data classes.Core) ShowAllMentorClass {
	return ShowAllMentorClass{
		ID:    data.ID,
		Image: data.Image,
		Name:  data.Name,
		Price: data.Price,
	}
}
