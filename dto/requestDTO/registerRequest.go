package requestdto

import (
	"net/mail"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
)

type RegisterRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

// Helper to build User Sruct from RegisterRequest

func BuildUser(r RegisterRequest) models.User {
	hasedPword, err := security.HashPword(r.Password)
	if err != nil {
		logger.Error("PasswordHash Eror: " + err.Error())
	}
	return models.User{
		Id:        "",
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
		Password:  hasedPword,
		CreatedAt: time.Now().Format(konstants.T_FORMAT),
		UpdatedAt: "",
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Request Validation function
func (req RegisterRequest) ValidateRequest() *ericerrors.EricError {
	if len(req.FirstName) < 2 || len(req.FirstName) > 20 {
		return ericerrors.New422Error("firstname char count must be 2 -20 ranged")
	}

	if len(req.LastName) < 2 {
		return ericerrors.New422Error("lastname char count must be 2 -20 ranged")
	}

	if !isValidEmail(req.Email) {
		return ericerrors.New422Error("Invalid Email Addrees")
	}

	if len(req.Phone) != 11 {
		return ericerrors.New422Error("phone char count must equal 11")
	}

	if len(req.Password) < 6 {
		return ericerrors.New422Error("password char count must be > 6")
	}

	return nil
}
