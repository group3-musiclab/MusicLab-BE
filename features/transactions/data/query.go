package data

import (
	"errors"
	"log"
	"musiclab-be/features/transactions"

	"gorm.io/gorm"
)

type transactionQuery struct {
	db *gorm.DB
}

// GetMentorTransaction implements transactions.TransactionData
func (*transactionQuery) GetMentorTransaction() {
	panic("unimplemented")
}

// GetStudentTransaction implements transactions.TransactionData
func (*transactionQuery) GetStudentTransaction() {
	panic("unimplemented")
}

// GetClass implements transactions.TransactionData

// GetMentorTransaction implements transactions.TransactionData

func New(db *gorm.DB) transactions.TransactionData {
	return &transactionQuery{
		db: db,
	}
}

// MakeTransaction implements transactions.TransactionData
func (tq *transactionQuery) MakeTransaction(newTransaction transactions.Core) error {
	cnv := CoreToData(newTransaction)

	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}
