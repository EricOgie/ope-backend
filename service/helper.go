package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/security"
)

//   ----------------------- PRIVATE METHOD ---------------------------- //
func getUserWithToken(user *models.User) (responsedto.OneUserDto, responsedto.OneUserDtoWithOtp) {
	// Gen OTP
	otp := security.GenerateOTP()
	// Construct UserDTOwithOTP from user
	userDTOWithOTP := user.ConvertToOneUserDtoWithOtp(otp)
	// Gen Token
	token := security.GenerateToken(userDTOWithOTP)
	// Contruct UserDTOwithToken
	userResponseDTOWithToken := user.ConvertToOneUserDto(token) //user.ConvertToOneUserDto(token)

	return userResponseDTOWithToken, userDTOWithOTP
}
