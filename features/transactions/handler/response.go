package handler

import (
	"musiclab-be/features/transactions"
)

type TransactionResponse struct {
	PaymentUrl string `json:"payment_url"`
}

func TransactionResponseResponse(data transactions.Core) TransactionResponse {
	return TransactionResponse{
		PaymentUrl: data.PaymentUrl,
	}
}

type ShowAllMentorTransaction struct {
	ID          uint    `json:"id"`
	StudentName string  `json:"student_name"`
	ClassName   string  `json:"class_name"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

func ShowAllMentorTransactionResponse(data transactions.Core) ShowAllMentorTransaction {
	StartDate := data.StartDate.Format("2006-01-02")
	EndDate := data.EndDate.Format("2006-01-02")
	return ShowAllMentorTransaction{
		ID:          data.ID,
		StudentName: data.Student.Name,
		ClassName:   data.Class.Name,
		StartDate:   StartDate,
		EndDate:     EndDate,
		Price:       data.Price,
		Status:      data.Status,
	}
}

type ShowAllStudentTransaction struct {
	ID         uint    `json:"id"`
	MentorName string  `json:"mentor_name"`
	ClassName  string  `json:"class_name"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
}

func ShowAllStudentTransactionResponse(data transactions.Core) ShowAllStudentTransaction {
	StartDate := data.StartDate.Format("2006-01-02")
	EndDate := data.EndDate.Format("2006-01-02")
	return ShowAllStudentTransaction{
		ID:         data.ID,
		MentorName: data.Mentor.Name,
		ClassName:  data.Class.Name,
		StartDate:  StartDate,
		EndDate:    EndDate,
		Price:      data.Price,
		Status:     data.Status,
	}
}
