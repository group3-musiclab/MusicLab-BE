package services

import (
	"errors"
	"log"
	"musiclab-be/features/classes"
	"musiclab-be/features/mentors"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
	"musiclab-be/utils/helper"
	"strings"
)

type transactionUseCase struct {
	qry        transactions.TransactionData
	qryClass   classes.ClassData
	qryMentor  mentors.MentorData
	qryStudent students.StudentData
}

func New(td transactions.TransactionData, md mentors.MentorData, sd students.StudentData, cd classes.ClassData) transactions.TransactionService {
	return &transactionUseCase{
		qry:        td,
		qryClass:   cd,
		qryMentor:  md,
		qryStudent: sd,
	}
}

// MakeTransaction implements transactions.TransactionService
func (tuc *transactionUseCase) MakeTransaction(newTransaction transactions.Core) (transactions.Core, error) {
	selectClass, errSelectClass := tuc.qryClass.GetMentorClassDetail(newTransaction.ClassID)
	if errSelectClass != nil {
		return transactions.Core{}, errSelectClass
	}

	selectStudent, errSelectStudent := tuc.qryStudent.SelectProfile(newTransaction.StudentID)
	if errSelectStudent != nil {
		return transactions.Core{}, errSelectStudent
	}

	midtransResponse, errSnap := helper.RequestSnapMidtrans(selectStudent, selectClass, newTransaction)
	if errSnap != nil {
		return transactions.Core{}, errSnap
	}

	countDay := int(selectClass.Duration * 30)

	endDate := newTransaction.StartDate.AddDate(0, 0, countDay)
	newTransaction.EndDate = endDate
	newTransaction.MentorID = selectClass.MentorID
	newTransaction.Price = selectClass.Price
	newTransaction.OrderID = midtransResponse.OrderID
	newTransaction.Status = "Pending"
	newTransaction.PaymentUrl = midtransResponse.PaymentUrl

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

	return midtransResponse, nil
}

func (tuc *transactionUseCase) GetMentorTransaction(mentorID uint) ([]transactions.Core, error) {
	res, err := tuc.qry.GetMentorTransaction(mentorID)

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
func (*transactionUseCase) GetStudentTransaction(studentID uint) ([]transactions.Core, error) {
	panic("unimplemented")
}
