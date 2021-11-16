package models

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id          string
	Firstname   string
	Lastname    string
	Email       string
	CreatedAt   string
	BankAccount map[string]interface{}
	Otp         int
	Wallet      map[string]interface{}
	Portfolio   interface{}
}

type PaymentClaim struct {
	TxRef         string
	Amount        string
	Currency      string
	PaymentType   string
	Mata          map[string]interface{}
	Customer      map[string]interface{}
	Customization map[string]interface{}
}

func RetrieveClaim(claim jwt.MapClaims) Claim {

	otpT, _ := strconv.Atoi(fmt.Sprintf("%v", claim["otp"]))
	return Claim{
		Id:          fmt.Sprintf("%v", claim["id"]),
		Firstname:   fmt.Sprintf("%v", claim["firstname"]),
		Lastname:    fmt.Sprintf("%v", claim["lastname"]),
		Email:       fmt.Sprintf("%v", claim["email"]),
		CreatedAt:   fmt.Sprintf("%v", claim["when"]),
		BankAccount: claim["bank_account"].(map[string]interface{}),
		Otp:         otpT,
		Wallet:      claim["wallet"].(map[string]interface{}),
	}
}

func RetrievePaymentClaim(claim jwt.MapClaims) PaymentClaim {
	return PaymentClaim{
		TxRef:         fmt.Sprintf("%v", claim["tx_ref"]),
		Amount:        fmt.Sprintf("%v", claim["amount"]),
		Currency:      fmt.Sprintf("%v", claim["currency"]),
		PaymentType:   fmt.Sprintf("%v", claim["payment_option"]),
		Mata:          claim["meta"].(map[string]interface{}),
		Customer:      claim["customer"].(map[string]interface{}),
		Customization: claim["customization"].(map[string]interface{}),
	}
}
