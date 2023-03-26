package handler

import (
	"musiclab-be/features/classes"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type classControll struct {
	srv classes.ClassService
}

// Delete implements classes.ClassHandler
func (*classControll) Delete() echo.HandlerFunc {
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
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "file size max 500kb"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success make a class",
		})
	}
}

func (cc *classControll) GetMentorClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("mentor_id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		res, err := cc.srv.GetMentorClass(uint(mentorID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowAllMentorClass{}
		for _, val := range res {
			result = append(result, ShowAllMentorClassResponse(val))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success show all mentor class",
		})
	}
}

func (cc *classControll) GetMentorClassDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("class_id")
		classID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}

		res, err := cc.srv.GetMentorClassDetail(uint(classID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ShowMentorClassDetailResponse(res),
			"message": "success show mentor class detail",
		})
	}
}

func (cc *classControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		id := c.Param("class_id")
		classID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		input := PostClass{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}
		updatedClass := addPostClassToCore(input)
		updatedClass.MentorID = mentorID

		checkFile, _, _ := c.Request().FormFile("image")
		if checkFile != nil {
			formHeader, err := c.FormFile("image")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Select a file to upload"})
			}
			input.FileHeader = *formHeader
		}

		res, err := cc.srv.Update(mentorID, uint(classID), input.FileHeader, updatedClass)

		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "file size max 500kb"})
			} else if strings.Contains(err.Error(), "validate") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "not registered") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "unable to process data"})
			}
		}

		result, err := ConvertClassUpdateResponse(res)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": err.Error(),
			})
		} else {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"data":    result,
				"message": "success edit class",
			})
		}
	}
}
