package requestdto

import (
	"net/mail"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
)

type RegisterRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserDetailsRequest struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AccountNo string `json:"account_no"`
	BankName  string `json:"bank_name"`
}

type BankAccountRequest struct {
	UserId    int    `json:"user_id"`
	AccountNo string `json:"account_no"`
	BankName  string `json:"bank_name"`
}

// Helper to build User Sruct from RegisterRequest
func BuildUser(r RegisterRequest) models.User {
	return models.User{
		Id:        "",
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
		Password:  r.Password,
		CreatedAt: time.Now().Format(konstants.T_FORMAT),
		UpdatedAt: "",
	}
}

func (u UserDetailsRequest) BuildQueryUser() models.QueryUser {
	return models.QueryUser{
		Id:          strconv.Itoa(u.Id),
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Phone:       u.Phone,
		AccountNo:   u.AccountNo,
		AccountName: u.BankName,
	}
}

func (b BankAccountRequest) BuildBankAccount() models.BankAccount {
	return models.BankAccount{
		UserId:        strconv.Itoa(b.UserId),
		AccountNumber: b.AccountNo,
		AccountName:   b.BankName,
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPword(pword string) bool {
	return len(pword) >= 6
}

// Request Validation function
func (req RegisterRequest) ValidateRequest() *ericerrors.EricError {
	if len(req.FirstName) < 2 || len(req.FirstName) > 20 {
		return ericerrors.New422Error(konstants.NAME_TOO_SHORT)
	}

	if len(req.LastName) < 2 {
		return ericerrors.New422Error(konstants.NAME_TOO_SHORT)
	}

	if !isValidEmail(req.Email) {
		return ericerrors.New422Error(konstants.INVALID_EMAIL)
	}

	if len(req.Phone) != 11 {
		return ericerrors.New422Error(konstants.PHONE_ERR)
	}

	if !isValidPword(req.Password) {
		return ericerrors.New422Error(konstants.INVALID_PWORD)
	}

	return nil
}
