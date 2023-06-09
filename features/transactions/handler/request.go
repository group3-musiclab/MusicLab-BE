package handler

import (
	"musiclab-be/features/transactions"
	"time"
)

type MakeTransaction struct {
	ClassID    uint    `json:"class_id" form:"class_id"`
	ScheduleID uint    `json:"schedule_id" form:"schedule_id"`
	StartDate  string  `json:"start_date" form:"start_date"`
	Price      float64 `json:"price" form:"price"`
}

func addMakeTransactionToCore(data MakeTransaction) transactions.Core {
	StartDate, _ := time.Parse("2006-01-02", data.StartDate)

	return transactions.Core{
		ClassID:    data.ClassID,
		ScheduleID: data.ScheduleID,
		StartDate:  StartDate,
		Price:      data.Price,
	}
}

type CheckTransactionRequest struct {
	OrderID string `json:"order_id" form:"order_id"`
	Status  string `json:"transaction_status" form:"transaction_status"`
}

func checkTransactionRequestToCore(data CheckTransactionRequest) transactions.Core {
	return transactions.Core{
		OrderID: data.OrderID,
		Status:  data.Status,
	}
}
