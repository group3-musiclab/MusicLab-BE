package handler

import (
	"musiclab-be/features/chats"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

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
		return c.JSON(http.StatusOK, helper.Response(consts.CHAT_SuccessInsert))
	}
}

// GetAll implements chats.ChatHandler
func (cc *chatControl) GetAll() echo.HandlerFunc {
	panic("unimplemented")
}

// GetByStudent implements chats.ChatHandler
func (cc *chatControl) GetByStudent() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv chats.ChatService) chats.ChatHandler {
	return &chatControl{
		srv: srv,
	}
}
