package security

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken takes responsedto.OneUserDto as aurg and return a string crtographed token
func GenerateToken(payload responsedto.OneUserDtoWithOtp) string {
	config := utils.LoadConfig(".")
	claim := genUserClaim(payload)
	return makeToken(claim, config)
}

// GenerateToken takes responsedto.OneUserDto as aurg and return a string crtographed token
func GeneTokenFromCompleteDTO(payload *responsedto.CompleteUserDTO) string {
	config := utils.LoadConfig(".")
	claim := genUserClaimFromCompleteUserDTO(payload)
	return makeToken(claim, config)
}

func GenPaymentToken(payload *responsedto.FlutterResponseDTO) string {
	config := utils.LoadConfig(".")
	claim := genPaymentClaim(payload)
	return makeToken(claim, config)
}

// Helpers func
func makeToken(claim jwt.Claims, config utils.Config) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(config.SigningKey))
	if err != nil {
		logger.Error("Failed To Sign Token. Error: " + err.Error())
		return ""
	} else {
		return signedToken
	}
}
