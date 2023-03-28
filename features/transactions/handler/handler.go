package handler

import (
	"musiclab-be/features/transactions"
	"musiclab-be/utils/consts"
	"musiclab-be/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionControll struct {
	srv transactions.TransactionService
}

func New(srv transactions.TransactionService) transactions.TransactionHandler {
	return &transactionControll{
		srv: srv,
	}
}

// MidtransNotification implements transactions.TransactionHandler
func (tc *transactionControll) MidtransNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := CheckTransactionRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		errUpdate := tc.srv.UpdateTransaction(checkTransactionRequestToCore(input))
		if errUpdate != nil {
			return c.JSON(helper.ErrorResponse(errUpdate))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get notification from midtrans",
		})
	}
}

// MakeTransaction implements transactions.TransactionHandler
func (tc *transactionControll) MakeTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		studentID := helper.ExtractTokenUserId(c)
		input := MakeTransaction{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format incorrect"})
		}

		transaction := addMakeTransactionToCore(input)
		transaction.StudentID = studentID

		res, err := tc.srv.MakeTransaction(transaction)

		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		result := TransactionResponseResponse(res)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success make transaction",
		})

	}

}

func (tc *transactionControll) GetMentorTransaction() echo.HandlerFunc {
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

		var limit int = 15
		limitParam := c.QueryParam("limit")
		if limitParam != "" {
			limitConv, errConv := strconv.Atoi(limitParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				limit = limitConv
			}
		}
		res, err := tc.srv.GetMentorTransaction(uint(mentorID), page, limit)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowAllMentorTransaction{}
		for _, val := range res {
			result = append(result, ShowAllMentorTransactionResponse(val))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success show mentor transaction history",
		})
	}
}

// GetStudentTransaction implements transactions.TransactionHandler
func (tc *transactionControll) GetStudentTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		studentID := helper.ExtractTokenUserId(c)
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

		var limit int = 15
		limitParam := c.QueryParam("limit")
		if limitParam != "" {
			limitConv, errConv := strconv.Atoi(limitParam)
			if errConv != nil {
				return c.JSON(http.StatusBadRequest, helper.Response(consts.HANDLER_InvalidPageParam))
			} else {
				limit = limitConv
			}
		}
		res, err := tc.srv.GetStudentTransaction(uint(studentID), page, limit)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}
		result := []ShowAllStudentTransaction{}
		for _, val := range res {
			result = append(result, ShowAllStudentTransactionResponse(val))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success show student transaction history",
		})
	}
}
