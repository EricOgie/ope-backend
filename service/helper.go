package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/security"
)

var purposeList = []string{"request", "login"}

func OTPNeeded(purpose string) bool {
	for _, a := range purposeList {
		if a == purpose {
			return true
		}
	}

	return false
}

func getDTOWithTokenAndOTP(user *models.CompleteUser, purpose string) responsedto.CompleteUserDTO {
	userDTO := user.ConvertToCompleteUserDTO()
	if !OTPNeeded(purpose) {
		userDTO.Otp = ""
	}
	userDTO.Otp = security.GenerateOTP()

	userDTO.Token = security.GeneTokenFromCompleteDTO(&userDTO)
	return userDTO
}
