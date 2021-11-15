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

func MakeClaim(claim jwt.MapClaims) Claim {

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
