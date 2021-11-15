package models

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

type User struct {
	Id          string `db:"id"`
	FirstName   string `json:"firstname" validate:"required,min=2,max=50" xml:"first_name"`
	LastName    string `json:"lastname" validate:"required,min=2,max=50" xml:"last_name"`
	Email       string `json:"email" validate:"email,required" xml:"email"`
	Phone       string `json:"phone" validate:"required" xml:"phone"`
	Password    string `json:"password" xml:"password" validate:"required,min=6"`
	AccountNo   string `db:"account_no" json:"account_no"`
	AccountName string `db:"account_name" json:"account_name"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type CompleteUser struct {
	Id          string `db:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `db:"created_at"`
	BankAccount BankAccount
	Wallet      Wallet
	Portfolio   []Stock
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmail struct {
	Email string
}

type EditResponse struct {
	Code    int
	Message string
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
	Create(User) (*CompleteUser, *ericerrors.EricError)
	VerifyUserAccount(VerifyUser) (*User, *ericerrors.EricError)
	Login(UserLogin) (*CompleteUser, *ericerrors.EricError)
	CompleteLogin(Claim) (*CompleteUser, *ericerrors.EricError)
	RequestPasswordChange(UserEmail) (*CompleteUser, *ericerrors.EricError)
	ChangePassword(UserLogin) (*responsedto.PlainResponseDTO, *ericerrors.EricError)
}

/**
* When serving user data to client side, it would be bad practice to send
* sensitive data like hashed user password alongside. Hence, data access object
* is used here
 */
// Getter function to conver User struct to UserDTO struc
func (user CompleteUser) ConvertToUserDto() responsedto.UserDto {
	return responsedto.UserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (user CompleteUser) ConvertToCompleteUserDTO() responsedto.CompleteUserDTO {
	return responsedto.CompleteUserDTO{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		BankAccount: responsedto.BankAccountDTO{AccountNo: user.BankAccount.AccountNumber, AccountName: user.BankAccount.AccountName},
		Wallet:      responsedto.WalletDTO{Amount: user.Wallet.Amount, Address: user.Wallet.Address},
		Token:       "",
		Portfolio:   user.Portfolio,
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

func (user User) ConvertToVeriyResponse(verified string) responsedto.VerifiedRESPONSE {
	return responsedto.VerifiedRESPONSE{
		Id:     user.Id,
		Email:  user.Email,
		Status: verified,
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

func (u UserLogin) GetPlainResponseDTO(code int, msg string) responsedto.PlainResponseDTO {
	return responsedto.PlainResponseDTO{
		Code:    code,
		Message: msg,
	}
}

// MakeAllInOneUserDTO function will output a complete user dTO with account, wallet and portfolio slice
func (qUser User) MakeCompleteUser(wallet *Wallet) CompleteUser {

	return CompleteUser{
		Id:          qUser.Id,
		FirstName:   qUser.FirstName,
		LastName:    qUser.LastName,
		Email:       qUser.Email,
		CreatedAt:   qUser.CreatedAt,
		BankAccount: BankAccount{AccountName: qUser.AccountName, AccountNumber: qUser.AccountNo},
		Wallet:      Wallet{Amount: wallet.Amount, Address: wallet.Address},
	}
}
