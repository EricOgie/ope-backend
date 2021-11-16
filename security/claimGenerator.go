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

func genPaymentClaim(p *responsedto.FlutterResponseDTO) jwt.Claims {
	return jwt.MapClaims{
		"tx_ref":         p.Tx_Ref,
		"amount":         p.Amount,
		"currency":       p.Currency,
		"payment_option": p.PaymentOption,
		"meta": map[string]string{
			"cusumer_id":   strconv.Itoa(p.Meta.ConsumerId),
			"consumer_mac": p.Meta.ConsumerMac,
		},
		"customer": map[string]string{
			"email": p.Customer.Name,
			"phone": p.Customer.PhoneNumber,
			"name":  p.Customer.Name,
		},
		"customization": map[string]string{
			"title": p.Customizations.Title,
			"desc":  p.Customizations.Description,
			"logo":  p.Customizations.Logo,
		},
	}
}
