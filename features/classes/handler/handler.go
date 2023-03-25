package handler

import (
	"log"
	"musiclab-be/features/classes"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type classControll struct {
	srv classes.ClassService
}

// Delete implements classes.ClassHandler
func (*classControll) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetMentorClass implements classes.ClassHandler
func (*classControll) GetMentorClass() echo.HandlerFunc {
	panic("unimplemented")
}

// GetMentorClassDetail implements classes.ClassHandler
func (*classControll) GetMentorClassDetail() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements classes.ClassHandler
func (*classControll) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv classes.ClassService) classes.ClassHandler {
	return &classControll{
		srv: srv,
	}
}

func (cc *classControll) PostClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		input := PostClass{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		mentorClass := addPostClassToCore(input)
		mentorClass.MentorID = mentorID

		checkFile, _, _ := c.Request().FormFile("image")
		if checkFile != nil {
			formHeader, err := c.FormFile("image")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}
		err = cc.srv.PostClass(input.FileHeader, mentorClass)
		if err != nil {
			log.Println("error running add mentor genre service: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "server problem"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success make a class",
		})
	}
}
