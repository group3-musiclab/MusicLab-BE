package data

import (
	"errors"
	"log"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/consts"

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

// SelectOne implements transactions.TransactionData
func (tq *transactionQuery) SelectOne(orderID string) (transactions.Core, error) {
	dataModel := Transaction{}
	txSelect := tq.db.Where("order_id = ?", orderID).First(&dataModel)
	if txSelect.Error != nil {
		return transactions.Core{}, errors.New(consts.QUERY_NotFound)
	}
	return (ToCore(dataModel)), nil
}

// UpdateTransaction implements transactions.TransactionData
func (tq *transactionQuery) UpdateTransaction(input transactions.Core) error {
	cnv := CoreToData(input)

	err := tq.db.Where("order_id = ?", cnv.OrderID).Updates(&cnv)
	if err != nil {
		log.Println("query error", err.Error)
		return errors.New("server error")
	}
	return nil
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

func (tq *transactionQuery) GetMentorTransaction(mentorID uint, limit, offset int) ([]transactions.Core, error) {
	res := []Transaction{}
	err := tq.db.Preload("Student").Preload("Class").Where("mentor_id = ?", mentorID).Limit(limit).Offset(offset).Find(&res).Error
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

func (tq *transactionQuery) GetStudentTransaction(studentID uint, limit, offset int) ([]transactions.Core, error) {
	res := []Transaction{}
	err := tq.db.Preload("Mentor").Preload("Class").Where("student_id = ?", studentID).Limit(limit).Offset(offset).Find(&res).Error
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
