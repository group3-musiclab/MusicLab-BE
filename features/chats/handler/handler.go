package handler

import (
	"musiclab-be/features/chats"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type chatControl struct {
	srv chats.ChatService
}

// Add implements chats.ChatHandler
func (cc *chatControl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := helper.ExtractTokenUserRole(c)
		input := AddChatRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		dataCore := addRequestToCore(input)
		dataCore.Role = role

		err := cc.srv.InsertChat(dataCore)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, helper.Response(consts.CHAT_SuccessInsert))
	}
}

// GetAll implements chats.ChatHandler
func (cc *chatControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		studentParam := c.QueryParam("student")
		studentID, errStudent := strconv.Atoi(studentParam)
		if errStudent != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(consts.HANDLER_InvalidIdStudent))
		}

		mentorParam := c.QueryParam("mentor")
		mentorID, errMentor := strconv.Atoi(mentorParam)
		if errMentor != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(consts.HANDLER_InvalidIdMentor))
		}

		dataCore, errSelect := cc.srv.GetAll(uint(mentorID), uint(studentID))
		if errSelect != nil {
			return c.JSON(helper.ErrorResponse(errSelect))
		}

		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.CHAT_SuccessAllChat, listCoreToAllChatResponse(dataCore)))
	}
}

// GetByStudent implements chats.ChatHandler
func (cc *chatControl) GetByStudent() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)

		var page int = 1
		pageParam := c.QueryParam("page")
		if pageParam != "" {
			pageConv, errConv := strconv.Atoi(pageParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				page = pageConv
			}
		}

		var limit int = 10
		limitParam := c.QueryParam("limit")
		if limitParam != "" {
			limitConv, errConv := strconv.Atoi(limitParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				limit = limitConv
			}
		}

		dataCore, errSelect := cc.srv.GetByStudent(page, limit, mentorID)
		if errSelect != nil {
			return c.JSON(helper.ErrorResponse(errSelect))
		}

		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.CHAT_SuccessChatByStudent, listCoreToChatByStudentResponse(dataCore)))
	}
}

func New(srv chats.ChatService) chats.ChatHandler {
	return &chatControl{
		srv: srv,
	}
}
