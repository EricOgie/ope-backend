package requestdto

import (
	"crypto/rand"
	"math/big"

	"github.com/EricOgie/ope-be/domain/models"
)

type PaymentFlutterRequest struct {
	Tx_Ref         string `json:"tx_ref"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	PaymentOptions string `json:"payment_options"`
	RedirectUrl    string `json:"redirect_url"`
	Meta           Meta
	Customer       Customer
	Customizations Customizations
}

type UserPayRequest struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	PaymentOptions string `json:"payment_options"`
	StockId        string `json:"stock_id"`
	StockSymbol    string `json:"stock_symbol"`
	StockImage     string `json:"stock_image"`
}

type Meta struct {
	ConsumerId string `json:"consumer_id"`
	StockImage string `json:"stock_image"`
}

type Customer struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Name        string `json:"name"`
}

type Customizations struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

func (userPay UserPayRequest) MakeFlutterPayRequest(claim models.Claim) PaymentFlutterRequest {
	return PaymentFlutterRequest{
		Tx_Ref:         claim.Firstname + "-tx-" + gencode() + userPay.StockSymbol,
		Amount:         userPay.Amount,
		Currency:       userPay.Currency,
		PaymentOptions: userPay.PaymentOptions,
		RedirectUrl:    "",
		Meta:           Meta{ConsumerId: claim.Id, StockImage: userPay.StockImage},
		Customer:       Customer{Email: claim.Email, PhoneNumber: "", Name: claim.Firstname},
		Customizations: Customizations{Title: "Stock Purchase", Description: "Buy stock", Logo: userPay.StockImage},
	}
}

func gencode() string {
	gen := ""
	for i := 0; i < 6; i++ {
		opeRand, _ := rand.Int(rand.Reader, big.NewInt(9))
		gen += opeRand.String()
	}

	return gen
}
