package tests

import (
	"net/http"
	"testing"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
)

func Test_profile_update_validation_returns_422_status_when_first_or_lastname_is_invalid(t *testing.T) {
	//Set-up
	req := requestdto.UserDetailsRequest{FirstName: "a", LastName: ""}
	//ACT
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong profile update Validation status code for invalid first or last name")
	}
}

func Test_profile_update_validation_returns_correct_error_msg_when_first_or_lastname_is_invalid(t *testing.T) {
	//Set-up
	req := requestdto.UserDetailsRequest{FirstName: "a", LastName: ""}
	//ACT
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.NAME_TOO_SHORT {
		t.Error("wrong profile update Validation error msg for invalid first or last name")
	}
}
func Test_profile_update_validation_returns_correct_msg_error_when_email_is_invalid(t *testing.T) {

	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie", Email: "ghdgfndb"}

	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.INVALID_EMAIL {
		t.Error("Wrong msg for profile email validation")
	}
}

func Test_profile_update_validation_returns_422_error_when_email_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie", Email: "ghdgfndb"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for profile email validation checker")
	}
}

func Test_profile_update_validation_returns_422_error_when_phone_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "7gdy64"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for profile validation when phone is invalid")
	}
}

func Test_profile_update_validation_returns_correct_error_msg_when_phone_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "7gdy64"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.PHONE_ERR {
		t.Error("wrong error msg for profile validation when phone is invalid")
	}
}

func Test_profile_update_validation_returns_422_error_when_Bank_acc_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "07053492875", AccountNo: "63h"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for profile validation when Acc no is invalid")
	}
}

func Test_profile_update_validation_returns_correct_error_msg_when_Acc_no_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "07053492875", AccountNo: "63h"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.ERR_ACC_NO {
		t.Error("wrong error msg for profile validation when Acc No is invalid")
	}
}

func Test_profile_update_validation_returns_422_error_when_bank_name_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "07053492875", AccountNo: "2085394463", BankName: "635g"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Code != http.StatusUnprocessableEntity {
		t.Error("wrong status code for profile validation when Bank name is invalid")
	}
}

func Test_profile_update_validation_returns_correct_error_msg_when_bank_name_is_invalid(t *testing.T) {
	// Set
	req := requestdto.UserDetailsRequest{FirstName: "Eric", LastName: "Ogie",
		Email: "sam@yahoo.com", Phone: "07053492875", AccountNo: "2085394463", BankName: "635g"}
	//ACC
	err := req.ValidateRequest()
	//Assert
	if err.Message != konstants.ERR_BANK_NAME {
		t.Error("wrong error msg for profile validation when bank name is invalid")
	}
}
