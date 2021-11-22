package requestdto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
)

type LoanPayRequest struct {
	UserId  string  `json:"user_id"`
	LoanId  string  `json:"loan_id"`
	Payment float64 `json:"payment"`
}

type LoanRequest struct {
	UserId   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Duration int     `json:"duration"`
}

func (req LoanPayRequest) ConvertToLoanPayment() models.LoanPayment {
	userId, _ := strconv.Atoi(req.UserId)
	loadId, _ := strconv.Atoi(req.LoanId)

	return models.LoanPayment{
		LoanId:    loadId,
		UserId:    userId,
		Payment:   req.Payment,
		CreatedAt: time.Now().Format(konstants.T_FORMAT),
	}
}

func (req LoanRequest) ConvertToLoan() models.Loan {
	userId, _ := strconv.Atoi(req.UserId)
	pkgAmount := interestFactor(req.Duration) * req.Amount
	pakageFloat := pkgAmount / float64(req.Duration)
	loanPackage := fmt.Sprintf("%f", pakageFloat) + " Per Month"
	duration := strconv.Itoa(int(req.Duration))
	return models.Loan{
		UserId:    userId,
		Amount:    req.Amount,
		Package:   loanPackage,
		Duration:  duration,
		CreatedAt: time.Now().Format(konstants.T_FORMAT),
	}
}

// -------------------------- VALIDATIONS ------------------------------- //

func (req LoanRequest) isValidAmount() bool {
	return req.Amount >= 2000.0
}

func (req LoanRequest) isValidDuration() bool {
	return req.Duration >= 6
}

func (req LoanRequest) Validate() *ericerrors.EricError {

	if !req.isValidAmount() {
		return ericerrors.New422Error(konstants.ERR_LOAN_AMT)
	}

	if !req.isValidDuration() {
		return ericerrors.New422Error(konstants.ERR_LOAN_DUR)
	}

	return nil
}

//
func (req LoanPayRequest) Validate() *ericerrors.EricError {

	if !req.isValidLoadId() {
		return ericerrors.New422Error(konstants.ERR_LOAN_ID)
	}

	if !req.isValidPayment() {
		return ericerrors.New422Error(konstants.ERR_PAY_AMT)
	}

	return nil
}

//
func (req LoanPayRequest) isValidLoadId() bool {
	_, err := strconv.Atoi(req.LoanId)
	return err == nil && req.LoanId != "0"
}

//
func (req LoanPayRequest) isValidPayment() bool {
	return req.Payment >= 500.0
}

func interestFactor(duration int) float64 {
	if duration == 6 {
		return 1.12
	} else if duration > 6 && duration < 8 {
		return 1.18
	} else {
		return 1.25
	}
}
