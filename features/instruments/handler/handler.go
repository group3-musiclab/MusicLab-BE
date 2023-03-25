package handler

import (
	"musiclab-be/features/instruments"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type instrumentControl struct {
	srv instruments.InstrumentService
}

// Add implements instruments.InstrumentHandler
func (ic *instrumentControl) Add() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements instruments.InstrumentHandler
func (ic *instrumentControl) Delete() echo.HandlerFunc {
	panic("unimplemented")
}

// GetAll implements instruments.InstrumentHandler
func (ic *instrumentControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ic.srv.SelectAll()
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		dataResponse := listCoreToResponse(data)
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.INSTRUMENT_SuccessSelectAll, dataResponse))
	}
}

// GetAllByMentorID implements instruments.InstrumentHandler
func (ic *instrumentControl) GetAllByMentorID() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv instruments.InstrumentService) instruments.InstrumentHandler {
	return &instrumentControl{
		srv: srv,
	}
}
