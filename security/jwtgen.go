package security

import (
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken takes responsedto.OneUserDto as aurg and return a string crtographed token
func GenerateToken(payload responsedto.OneUserDtoWithOtp) string {
	config, _ := utils.LoadConfig(".")
	claim := genUserClaim(payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(config.SigningKey))
	// Handle token error siginig
	if err != nil {
		logger.Error("Failed To Sign Token. Error: " + err.Error())
		return ""
	} else {
		return signedToken
	}
}

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

func genUserClaimFromCompleteUser(payload models.CompleteUser) jwt.MapClaims {
	return jwt.MapClaims{
		"id":        payload.Id,
		"firstname": payload.FirstName,
		"lastname":  payload.LastName,
		"email":     payload.Email,
		"when":      payload.CreatedAt,
		"exp":       time.Now().Add(time.Duration(konstants.EXP_TIME)).Unix(),
		"portfolio": payload.Portfolio,
	}
}

// GenerateToken takes responsedto.OneUserDto as aurg and return a string crtographed token
func GeneTokenFromCompleteDTO(payload models.CompleteUser) string {
	config, _ := utils.LoadConfig(".")
	claim := genUserClaimFromCompleteUser(payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(config.SigningKey))
	// Handle token error siginig
	if err != nil {
		logger.Error("Failed To Sign Token. Error: " + err.Error())
		return ""
	} else {
		return signedToken
	}
}
