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

func (tc *transactionControll) GetMentorTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		mentorID := helper.ExtractTokenUserId(c)
		res, err := tc.srv.GetMentorTransaction(uint(mentorID))
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
		res, err := tc.srv.GetStudentTransaction(uint(studentID))
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
