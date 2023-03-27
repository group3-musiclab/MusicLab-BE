package handler

import (
	"musiclab-be/features/transactions"
	"musiclab-be/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type transactionControll struct {
	srv transactions.TransactionService
}

// GetMentorTransaction implements transactions.TransactionHandler
func (*transactionControll) GetMentorTransaction() echo.HandlerFunc {
	panic("unimplemented")
}

// GetStudentTransaction implements transactions.TransactionHandler
func (*transactionControll) GetStudentTransaction() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv transactions.TransactionService) transactions.TransactionHandler {
	return &transactionControll{
		srv: srv,
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
			"message": "success make a class",
		})

	}

}
