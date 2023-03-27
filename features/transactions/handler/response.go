package handler

import "musiclab-be/features/transactions"

type TransactionResponse struct {
	PaymentUrl string `json:"payment_url"`
}

func TransactionResponseResponse(data transactions.Core) TransactionResponse {
	return TransactionResponse{
		PaymentUrl: data.PaymentUrl,
	}
}
