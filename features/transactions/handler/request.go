package handler

import "musiclab-be/features/transactions"

type MakeTransaction struct {
	ClassID    uint `json:"class_id" form:"class_id"`
	ScheduleID uint `json:"schedule_id" form:"schedule_id"`
}

func addMakeTransactionToCore(data MakeTransaction) transactions.Core {
	return transactions.Core{
		ClassID:    data.ClassID,
		ScheduleID: data.ScheduleID,
	}
}
