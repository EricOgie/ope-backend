package models

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

type User struct {
	Id        string `db:"id"`
	FirstName string `json:"firstname" validate:"required,min=2,max=50" xml:"first_name"`
	LastName  string `json:"lastname" validate:"required,min=2,max=50" xml:"last_name"`
	Email     string `json:"email" validate:"email,required" xml:"email"`
	Phone     string `json:"phone" validate:"required" xml:"phone"`
	Password  string `json:"password" xml:"password" validate:"required,min=6"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type CompleteUser struct {
	Id        string `db:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `db:"created_at"`
	Portfolio []Stock
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmail struct {
	Email string
}

type VerifyUser struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt string
}

type VerifyUserResponse struct {
	Id    string `roken:"email"`
	Email string
	Token string `roken:"email"`
}

// Add User adapter port
type UserRepositoryPort interface {
	FindAll() (*[]responsedto.UserDto, *ericerrors.EricError)
	Create(User) (*User, *ericerrors.EricError)
	VerifyUserAccount(VerifyUser) (*User, *ericerrors.EricError)
	Login(UserLogin) (*User, *ericerrors.EricError)
	CompleteLogin(Claim) (*CompleteUser, *ericerrors.EricError)
	// RequestPasswordChange(UserEmail) (*User, *ericerrors.EricError)
}

/**
* When serving user data to client side, it would be bad practice to send
* sensitive data like hashed user password alongside. Hence, data access object
* is used here
 */
// Getter function to conver User struct to UserDTO struc
func (user User) ConvertToUserDto() responsedto.UserDto {
	return responsedto.UserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user User) ConvertToOneUserDto(token string) responsedto.OneUserDto {
	return responsedto.OneUserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}
}

func (user User) ConvertToOneUserDtoWithOtp(otp int) responsedto.OneUserDtoWithOtp {
	return responsedto.OneUserDtoWithOtp{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		OTP:       otp,
		CreatedAt: user.CreatedAt,
	}
}

func (user User) ConvertToVeriyResponse() responsedto.VerifiedRESPONSE {
	return responsedto.VerifiedRESPONSE{
		Id:     user.Id,
		Email:  user.Email,
		Status: "Verified",
	}
}

func (user CompleteUser) ConvertToCompleteUserDTO(tokenString string) responsedto.CompleteUserDTO {
	return responsedto.CompleteUserDTO{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     tokenString,
		Portfolio: user.Portfolio,
	}
}

func (v VerifyUser) GetUserFromVerify() User {
	return User{
		Id:        v.Id,
		FirstName: v.FirstName,
		LastName:  v.LastName,
		Email:     v.Email,
		CreatedAt: v.CreatedAt,
	}
}
