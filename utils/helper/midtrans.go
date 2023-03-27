package helper

import (
	"musiclab-be/app/config"
	"musiclab-be/features/classes"
	"musiclab-be/features/students"
	"musiclab-be/features/transactions"
	"strconv"

	"github.com/lithammer/shortuuid/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type CheckTransaction struct {
	OrderID string `json:"order_id" form:"order_id"`
	Status  string `json:"transaction_status" form:"transaction_status"`
}

func RequestSnapMidtrans(student students.Core, class classes.Core, input transactions.Core) (transactions.Core, error) {
	// request midtrans snap
	var snapClient = snap.Client{}
	snapClient.New(config.SERVER_KEY_MIDTRANS, midtrans.Sandbox)

	// parsing student id, class id, uuid
	student_id := strconv.Itoa(int(input.StudentID))
	class_id := strconv.Itoa(int(input.ClassID))
	uuid := shortuuid.New()
	orderID := "ALTA-MusicLab-" + student_id + "-" + uuid

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       student.Name,
		Phone:       student.Phone,
		Address:     student.Address,
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(class.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    student.Name,
			Email:    student.Email,
			Phone:    student.Phone,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Class-" + class_id,
				Qty:   int32(1),
				Price: int64(class.Price),
				Name:  class.Name,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		return transactions.Core{}, errSnap
	}

	midtransResponse := transactions.Core{
		PaymentUrl: response.RedirectURL,
		OrderID:    orderID,
	}

	return midtransResponse, nil
}

func CallBackMidtrans(orderID string) (CheckTransaction, error) {
	var coreAPIClient = coreapi.Client{}
	coreAPIClient.New(config.SERVER_KEY_MIDTRANS, midtrans.Sandbox)

	response, err := coreapi.CheckTransaction(orderID)
	if err != nil {
		return CheckTransaction{}, err
	}

	transactionStatus := CheckTransaction{
		OrderID: response.OrderID,
		Status:  response.TransactionStatus,
	}

	return transactionStatus, nil
}
