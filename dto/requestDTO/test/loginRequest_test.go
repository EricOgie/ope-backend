package requestdto

import (
	"net/http"
	"testing"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
)

func Test_login_email_validation_returns_422_error_when_email_is_invalid(t *testing.T) {
	//Set-up
	req := requestdto.LoginRequest{Email: "dhsbshd"}
	//ACT
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for login email validation checker")
	}
}

func Test_login_email_validation_returns_correct_msg_when_email_is_invalid(t *testing.T) {
	// Set
	req := requestdto.LoginRequest{Email: "dhsbshd"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.INVALID_EMAIL {
		t.Error("Wrong msg for login email validation")
	}
}

func Test_login_password_validation_returns_correct_msg_when_password_is_invalid(t *testing.T) {
	// Set
	req := requestdto.LoginRequest{Email: "ericogia@yahoo.com", Password: ""}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.INVALID_PWORD {
		t.Error("Wrong msg for login password validation")
	}
}
