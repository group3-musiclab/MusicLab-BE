package handler

import (
	"musiclab-be/features/chats"

	"github.com/labstack/echo/v4"
)

type chatControl struct {
	srv chats.ChatService
}

// Add implements chats.ChatHandler
func (cc *chatControl) Add() echo.HandlerFunc {
	panic("unimplemented")
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
