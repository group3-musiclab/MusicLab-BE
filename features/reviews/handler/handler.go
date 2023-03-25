package handler

import (
	"log"
	"musiclab-be/features/reviews"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type reviewControll struct {
	srv reviews.ReviewService
}

func New(srv reviews.ReviewService) reviews.ReviewHandler {
	return &reviewControll{
		srv: srv,
	}
}

// PostMentorReview implements reviews.ReviewHandler
func (rc *reviewControll) PostMentorReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("mentor_id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		studentID := helper.ExtractTokenUserId(c)
		idMentor := helper.ExtractTokenUserId(c)
		input := AddMentorReview{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		mentorGenre := addMentorReviewToCore(input)
		mentorGenre.StudentID = studentID
		mentorGenre.MentorID = idMentor
		err = rc.srv.PostMentorReview(uint(mentorID), mentorGenre)
		if err != nil {
			log.Println("error running add mentor genre service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success make a review",
		})
	}
}

func (rc *reviewControll) GetMentorReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("mentor_id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		res, err := rc.srv.GetMentorReview(uint(mentorID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowAllMentorReview{}
		for _, val := range res {
			result = append(result, ShowAllMentorReviewResponse(val))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success make a review",
		})
	}
}
