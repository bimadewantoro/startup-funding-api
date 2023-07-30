package payment

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"startup-funding/user"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	var client snap.Client

	midtrans.ServerKey = "SB-Mid-server-x85FfyzK9-oRIGsjTlLV2QGg"
	midtrans.Environment = midtrans.Sandbox

	client.New(midtrans.ServerKey, midtrans.Environment)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	resp, err := client.CreateTransactionUrl(req)
	if err != nil {
		return "", err
	}

	return resp, nil
}
