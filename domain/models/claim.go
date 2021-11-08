package models

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id         string
	Firstname  string
	Lastname   string
	Email      string
	Created_at string
	When       string
	Otp        int
}

func MakeClaim(claim jwt.MapClaims) Claim {
	otpT, _ := strconv.Atoi(fmt.Sprintf("%v", claim["otp"]))
	return Claim{
		Id:        fmt.Sprintf("%v", claim["id"]),
		Firstname: fmt.Sprintf("%v", claim["firstname"]),
		Lastname:  fmt.Sprintf("%v", claim["lastname"]),
		Email:     fmt.Sprintf("%v", claim["email"]),
		When:      fmt.Sprintf("%v", claim["when"]),
		Otp:       otpT,
	}
}
