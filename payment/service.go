package payment

import (
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"log"
	"os"
	"startup-funding/user"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	var client snap.Client

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	environment := os.Getenv("MIDTRANS_ENVIRONMENT")

	midtrans.ServerKey = serverKey

	switch environment {
	case "sandbox":
		midtrans.Environment = midtrans.Sandbox
	case "production":
		midtrans.Environment = midtrans.Production
	default:
		log.Fatal("Invalid MIDTRANS_ENVIRONMENT value in .env")
	}

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
