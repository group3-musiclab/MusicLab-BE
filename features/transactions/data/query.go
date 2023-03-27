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

func New(db *gorm.DB) transactions.TransactionData {
	return &transactionQuery{
		db: db,
	}
}

func (tq *transactionQuery) MakeTransaction(newTransaction transactions.Core) error {
	cnv := CoreToData(newTransaction)

	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("server error")
	}
	return nil
}

func (tq *transactionQuery) GetMentorTransaction(mentorID uint) ([]transactions.Core, error) {
	res := []Transaction{}
	err := tq.db.Preload("Student").Preload("Class").Where("mentor_id = ?", mentorID).Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []transactions.Core{}, errors.New("server error")
	}

	result := []transactions.Core{}
	for _, val := range res {
		result = append(result, ToCore(val))
	}

	return result, nil
}

func (*transactionQuery) GetStudentTransaction(studentID uint) ([]transactions.Core, error) {
	panic("unimplemented")
}
