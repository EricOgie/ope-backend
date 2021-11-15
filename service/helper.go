package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/security"
)

//   ----------------------- PRIVATE METHOD ---------------------------- //
// func getUserWithToken(user *models.CompleteUser) (responsedto.OneUserDto, responsedto.OneUserDtoWithOtp) {
// 	// Gen OTP
// 	otp := security.GenerateOTP()
// 	// Construct UserDTOwithOTP from user
// 	userDTOWithOTP := user.ConvertToOneUserDtoWithOtp(otp)
// 	// Gen Token
// 	token := security.GenerateToken(userDTOWithOTP)
// 	// Contruct UserDTOwithToken
// 	userResponseDTOWithToken := user.ConvertToOneUserDto(token) //user.ConvertToOneUserDto(token)

// 	return userResponseDTOWithToken, userDTOWithOTP
// }
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
