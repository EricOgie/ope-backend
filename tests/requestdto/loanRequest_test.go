package tests

import (
	"net/http"
	"testing"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
)

func Test_loan_request_validation_spills_422_status_when_loan_amount_is_less_than_2000(t *testing.T) {

	req := requestdto.LoanRequest{Amount: 1200}

	err := req.Validate()

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Loan amount validation code error")
	}
}

func Test_loan_request_validation_spills_correct_msg_when_loan_amount_is_less_than_2000(t *testing.T) {

	req := requestdto.LoanRequest{Amount: 1200}

	err := req.Validate()

	if err.Message != konstants.ERR_LOAN_AMT {
		t.Error("Loan amount validation msg error")
	}
}

func Test_loan_request_validation_spills_correct_msg_when_loan_duration_is_not_btw_5_and_13_months(t *testing.T) {

	req := requestdto.LoanRequest{Amount: 1200, Duration: 2}

	err := req.Validate()

	if err.Message != konstants.ERR_LOAN_AMT {
		t.Error("Loan amount validation msg error")
	}
}

func Test_loan_request_validation_spills_422_status_when_loan_duration_is_not_btw_5_and_13_months(t *testing.T) {

	req := requestdto.LoanRequest{Amount: 1200, Duration: 2}

	err := req.Validate()

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Loan duration validation code error")
	}
}

// LOan Repayment Request test
