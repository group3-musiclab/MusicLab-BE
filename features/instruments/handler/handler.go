package handler

import (
	"musiclab-be/features/instruments"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type instrumentControl struct {
	srv instruments.InstrumentService
}

// Add implements instruments.InstrumentHandler
func (ic *instrumentControl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := helper.ExtractTokenUserId(c)
		input := MentorInstrumentRequest{}
		errBind := c.Bind(&input)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.AUTH_ErrorBind))
		}

		dataCore := insertRequestToCore(input)
		dataCore.MentorID = id

		err := ic.srv.Insert(dataCore)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, helper.Response(consts.INSTRUMENT_SuccessInsert))
	}
}

// Delete implements instruments.InstrumentHandler
func (ic *instrumentControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		instrumentID := c.Param("id")
		idConv, errConv := strconv.Atoi(instrumentID)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}

		err := ic.srv.Delete(mentorID, uint(idConv))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, helper.Response(consts.INSTRUMENT_SuccessDelete))
	}
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
	return func(c echo.Context) error {
		id := c.Param("id")
		mentorID, errConv := strconv.Atoi(id)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_ErrorIdParam))
		}
		data, err := ic.srv.SelectAllByMentorID(uint(mentorID))
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		dataResponse := listCoreToMentorInstrumentResponse(data)
		return c.JSON(http.StatusOK, helper.ResponseWithData(consts.INSTRUMENT_SuccessSelectAll, dataResponse))
	}
}

func New(srv instruments.InstrumentService) instruments.InstrumentHandler {
	return &instrumentControl{
		srv: srv,
	}
}
