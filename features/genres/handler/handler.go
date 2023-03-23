package handler

import (
	"log"
	"musiclab-be/features/genres"
	"net/http"

	"github.com/labstack/echo/v4"
)

type genreControll struct {
	srv genres.GenreService
}

func New(srv genres.GenreService) genres.GenreHandler {
	return &genreControll{
		srv: srv,
	}
}

func (gc *genreControll) AddMentorGenre() echo.HandlerFunc {
	return func(c echo.Context) error {
		// mentorID := helper.ExtractTokenUserId(e)
		input := AddMentorGenre{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		res, err := gc.srv.AddMentorGenre(input.GenreID, *ReqToCore(input))
		if err != nil {
			log.Println("error running add mentor genre service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success add mentor genre",
		})
	}
}

// Delete implements genres.GenreHandler
func (*genreControll) Delete(e echo.Context) echo.HandlerFunc {
	panic("unimplemented")
}

// GetGenre implements genres.GenreHandler
func (*genreControll) GetGenre(e echo.Context) echo.HandlerFunc {
	panic("unimplemented")
}

// GetMentorGenre implements genres.GenreHandler
func (*genreControll) GetMentorGenre(e echo.Context) echo.HandlerFunc {
	panic("unimplemented")
}
