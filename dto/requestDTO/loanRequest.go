package requestdto

import (
	"fmt"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
)

type LoanPayRequest struct {
	UserId  string  `json:"user_id"`
	LoanId  string  `json:"loan_id"`
	Payment float64 `json:"payment"`
}

type LoanRequest struct {
	UserId   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Duration float64 `json:"duration"`
}

func (req LoanPayRequest) ConvertToLoanPayment() models.LoanPayment {
	userId, _ := strconv.Atoi(req.UserId)
	loadId, _ := strconv.Atoi(req.LoanId)

	return models.LoanPayment{
		LoanId:    loadId,
		UserId:    userId,
		Payment:   req.Payment,
		CreatedAt: time.Now().String(),
	}
}

func (req LoanRequest) ConvertToLoan() models.Loan {
	userId, _ := strconv.Atoi(req.UserId)
	pakageFloat := (req.Amount / req.Duration)
	loanPackage := fmt.Sprintf("%f", pakageFloat) + " Per Month"
	duration := fmt.Sprintf("%f", req.Duration)
	return models.Loan{
		UserId:    userId,
		Amount:    req.Amount,
		Package:   loanPackage,
		Duration:  duration,
		CreatedAt: time.Now().String(),
	}
}

// -------------------------- VALIDATIONS ------------------------------- //

func (req LoanRequest) isValidAmount() bool {
	return req.Amount > 10000.0
}

func (req LoanRequest) isValidDuration() bool {
	return req.Duration >= 6
}

func (req LoanRequest) Validate() *ericerrors.EricError {

	if !req.isValidAmount() {
		return ericerrors.New422Error("Invalid Loan Amount")
	}

	if !req.isValidDuration() {
		return ericerrors.New422Error("Invalid Loan Duration")
	}

	return nil
}

//
func (req LoanPayRequest) Validate() *ericerrors.EricError {

	if !req.isValidLoadId() {
		return ericerrors.New422Error("Invalid Laon Id")
	}

	if !req.isValidPayment() {
		return ericerrors.New422Error("Invalid Payment Amount")
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
