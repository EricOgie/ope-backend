package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Id         string
	Firstname  string
	Lastname   string
	Email      string
	Created_at string
	When       string
}

func makeClaim(claim jwt.MapClaims) Claim {
	return Claim{
		Id:        fmt.Sprintf("%v", claim["id"]),
		Firstname: fmt.Sprintf("%v", claim["firstname"]),
		Lastname:  fmt.Sprintf("%v", claim["lastname"]),
		Email:     fmt.Sprintf("%v", claim["email"]),
		When:      fmt.Sprintf("%v", claim["when"]),
	}
}
