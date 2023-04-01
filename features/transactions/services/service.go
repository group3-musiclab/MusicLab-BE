package services

import (
	"errors"
	"musiclab-be/features/classes"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/helper"
	"strings"
	"time"
)

type transactionUseCase struct {
	qry        transactions.TransactionData
	qryClass   classes.ClassData
	qryStudent students.StudentData
}

func New(td transactions.TransactionData, sd students.StudentData, cd classes.ClassData) transactions.TransactionService {
	return &transactionUseCase{
		qry:        td,
		qryClass:   cd,
		qryStudent: sd,
	}
}

// UpdateTransaction implements transactions.TransactionService
func (tuc *transactionUseCase) UpdateTransaction(input transactions.Core) error {
	err := tuc.qry.UpdateTransaction(input)
	if err != nil {
		return err
	}
	return nil
}

func (tuc *transactionUseCase) MakeTransaction(newTransaction transactions.Core) (transactions.Core, error) {
	// start date validation
	year := time.Now().Year()
	month := time.Now().Month()
	date := time.Now().Day()
	today := time.Date(year, month, date, 0, 0, 0, 0, time.UTC)

	dateValidation := newTransaction.StartDate.Before(today)
	if dateValidation {
		return transactions.Core{}, errors.New("minimum start date input is today")
	}

	selectClass, errSelectClass := tuc.qryClass.GetMentorClassDetail(newTransaction.ClassID)
	if errSelectClass != nil {
		return transactions.Core{}, errSelectClass
	}

	selectStudent, errSelectStudent := tuc.qryStudent.SelectProfile(newTransaction.StudentID)
	if errSelectStudent != nil {
		return transactions.Core{}, errSelectStudent
	}

	countDay := int(selectClass.Duration * 30)

	endDate := newTransaction.StartDate.AddDate(0, 0, countDay)
	newTransaction.EndDate = endDate
	newTransaction.MentorID = selectClass.MentorID
	newTransaction.Price = selectClass.Price
	newTransaction.Status = "Pending"

	midtransResponse, errSnap := helper.RequestSnapMidtrans(selectStudent, selectClass, newTransaction)
	if errSnap != nil {
		return transactions.Core{}, errSnap
	}

	newTransaction.OrderID = midtransResponse.OrderID
	newTransaction.PaymentUrl = midtransResponse.PaymentUrl

	err := tuc.qry.MakeTransaction(newTransaction)
	if err != nil {
		return transactions.Core{}, err
	}

	return midtransResponse, nil
}

func (tuc *transactionUseCase) GetMentorTransaction(mentorID uint, page, limit int) ([]transactions.Core, error) {
	offset := (page - 1) * limit
	res, err := tuc.qry.GetMentorTransaction(mentorID, limit, offset)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "review not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []transactions.Core{}, errors.New(msg)
	}

	return res, nil
}

// GetStudentTransaction implements transactions.TransactionService
func (tuc *transactionUseCase) GetStudentTransaction(studentID uint, page, limit int) ([]transactions.Core, error) {
	offset := (page - 1) * limit
	res, err := tuc.qry.GetStudentTransaction(studentID, limit, offset)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "review not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []transactions.Core{}, errors.New(msg)
	}

	return res, nil
}
