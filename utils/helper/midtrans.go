package helper

import (
	"musiclab-be/app/config"
	"musiclab-be/features/transactions"
	"strconv"

	"github.com/lithammer/shortuuid/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func RequestSnapMidtrans(student transactions.Core, class transactions.Core, input transactions.Core) (transactions.Core, error) {
	// request midtrans snap
	var snapClient = snap.Client{}
	snapClient.New(config.SERVER_KEY_MIDTRANS, midtrans.Sandbox)

	// parsing student id, class id, uuid
	student_id := strconv.Itoa(int(input.StudentID))
	class_id := strconv.Itoa(int(input.ClassID))
	uuid := shortuuid.New()

	// customer
	custAddress := &midtrans.CustomerAddress{
		FName:       student.Student.Name,
		Phone:       student.Student.Phone,
		Address:     student.Student.Address,
		CountryCode: "IDN",
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ALTA-MusicLab-" + student_id + "-" + uuid,
			GrossAmt: int64(class.Class.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    student.Student.Name,
			Email:    student.Student.Email,
			Phone:    student.Student.Phone,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Class-" + class_id,
				Qty:   int32(class.Class.Qty),
				Price: int64(class.Class.Price),
				Name:  class.Class.Name,
			},
		},
	}

	response, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		return transactions.Core{}, errSnap
	}

	midtransResponse := transactions.Core{
		PaymentUrl: response.RedirectURL,
	}

	return midtransResponse, nil
}
