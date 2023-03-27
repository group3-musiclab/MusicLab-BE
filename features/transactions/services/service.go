package services

import (
	"errors"
	"log"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/helper"
	"strings"
)

type transactionUseCase struct {
	qry transactions.TransactionData
}

// GetMentorTransaction implements transactions.TransactionService
func (*transactionUseCase) GetMentorTransaction() {
	panic("unimplemented")
}

// GetStudentTransaction implements transactions.TransactionService
func (*transactionUseCase) GetStudentTransaction() {
	panic("unimplemented")
}

func New(td transactions.TransactionData) transactions.TransactionService {
	return &transactionUseCase{
		qry: td,
	}
}

// MakeTransaction implements transactions.TransactionService
func (tuc *transactionUseCase) MakeTransaction(newTransaction transactions.Core) (transactions.Core, error) {
	err := tuc.qry.MakeTransaction(newTransaction)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return transactions.Core{}, errors.New(msg)
	}

	midtransResponse, errSnap := helper.RequestSnapMidtrans(newTransaction)
	if errSnap != nil {
		return transactions.Core{}, errSnap
	}
	return midtransResponse, nil
}
