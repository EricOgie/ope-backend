package tests

import (
	"net/http"
	"testing"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
)

// This will test validation for for each field of register request
func Test_name_validation_returns_422_error_when_firstname_or_lastname_is_invalid(t *testing.T) {
	//Set-up
	req := requestdto.RegisterRequest{FirstName: "a", LastName: ""}
	//ACT
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for name validation checker")
	}
}

func Test_register_name_validation_returns_correct_error_message_when_firstname_or_lastname_is_invalid(t *testing.T) {
	//Set-up
	req := requestdto.RegisterRequest{FirstName: "a", LastName: ""}
	//ACT
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.NAME_TOO_SHORT {
		t.Error("wrong name Validation error msg")
	}
}

func Test_register_email_validation_returns_422_status_code_when_email_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes", Email: "didi"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Wrong status code for email validation")
	}
}

func Test_register_email_validation_returns_correct_msg_when_email_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes", Email: "didi"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.INVALID_EMAIL {
		t.Error("Wrong register email msg for email validation")
	}
}

func Test_register_phone_validation_returns_422_status_code_when_phone_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes",
		Email: "ericogia@yahoo.com", Phone: ""}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Wrong status code for phone validation")
	}
}

func Test_register_phone_validation_returns_correct_msg_when_phone_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes",
		Email: "ericogia@yahoo.com", Phone: ""}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.PHONE_ERR {
		t.Error("Wrong msg for phone validation")
	}
}

func Test_register_password_validation_returns_422_status_when_password_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes",
		Email: "ericogia@yahoo.com", Phone: "07053492875", Password: "3er"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Wrong msg for password validation")
	}
}
func Test_register_password_validation_returns_correct_msg_when_password_entered_is_of_invalid_format(t *testing.T) {
	// Set
	req := requestdto.RegisterRequest{FirstName: "James", LastName: "janes",
		Email: "ericogia@yahoo.com", Phone: "07053492875", Password: "3er"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.INVALID_PWORD {
		t.Error("Wrong msg for password validation")
	}
}
