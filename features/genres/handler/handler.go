package handler

import (
	"log"
	"musiclab-be/features/genres"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

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
		mentorID := helper.ExtractTokenUserId(c)
		input := AddMentorGenre{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		mentorGenre := addMentorGenreToCore(input)
		mentorGenre.MentorID = mentorID
		res, err := gc.srv.AddMentorGenre(mentorGenre)
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

func (gc *genreControll) GetGenre() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := gc.srv.GetGenre()
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		result := []ShowAllGenre{}
		for _, val := range res {
			result = append(result, ShowAllGenreResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all genres",
		})
	}
}

// GetMentorGenre implements genres.GenreHandler
func (gc *genreControll) GetMentorGenre() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		res, err := gc.srv.GetMentorGenre(uint(mentorID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowAllMentorGenre{}
		for _, val := range res {
			result = append(result, ShowAllMentorGenreResponse(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show all genres",
		})
	}
}

func (gc *genreControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		paramID := c.Param("genre_id")
		genreID, _ := strconv.Atoi(paramID)
		err := gc.srv.Delete(mentorID, uint(genreID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete mentor genre",
		})
	}
}
