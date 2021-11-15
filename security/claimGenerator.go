package security

import (
	"fmt"
	"strconv"
	"time"

	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/dgrijalva/jwt-go"
)

func genUserClaim(payload responsedto.OneUserDtoWithOtp) jwt.MapClaims {
	logger.Info("fircst gen")
	return jwt.MapClaims{
		"id":        payload.Id,
		"firstname": payload.FirstName,
		"lastname":  payload.LastName,
		"email":     payload.Email,
		"otp":       payload.OTP,
		"when":      payload.CreatedAt,
		"exp":       time.Now().Add(time.Duration(konstants.EXP_TIME)).Unix(),
	}
}

func genUserClaimFromCompleteUserDTO(payload *responsedto.CompleteUserDTO) jwt.MapClaims {
	logger.Info("2nd gen")
	otp, err := strconv.Atoi(payload.Otp)
	if err != nil {
		logger.Error("CONVERSION ERR")
	}
	return jwt.MapClaims{
		"id":        payload.Id,
		"firstname": payload.FirstName,
		"lastname":  payload.LastName,
		"email":     payload.Email,
		"when":      payload.CreatedAt,
		"exp":       time.Now().Add(time.Duration(konstants.EXP_TIME)).Unix(),
		"bank_account": map[string]string{
			"account_no":   payload.BankAccount.AccountNo,
			"account_name": payload.BankAccount.AccountName,
		},
		"otp": otp,
		"wallet": map[string]string{
			"address": payload.Wallet.Address,
			"amount":  fmt.Sprintf("%f", payload.Wallet.Amount),
		},
		"portfolio": payload.Portfolio,
	}
}
