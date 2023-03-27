package handler

import "musiclab-be/features/transactions"

type TransactionResponse struct {
	ID         uint   `json:"id"`
	ClassID    uint   `json:"class_id"`
	PaymentUrl string `json:"payment_url"`
}

func TransactionResponseResponse(data transactions.Core) TransactionResponse {
	return TransactionResponse{
		ID:         data.ID,
		ClassID:    data.ClassID,
		PaymentUrl: data.PaymentUrl,
	}
}
