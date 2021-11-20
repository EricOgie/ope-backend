package requestdto

import (
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
	Package  string  `json:"package"`
	Duration string  `json:"duration"`
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
	return models.Loan{
		UserId:    userId,
		Amount:    req.Amount,
		Package:   req.Package + " Per Month",
		Duration:  req.Duration,
		CreatedAt: time.Now().String(),
	}
}

// -------------------------- VALIDATIONS ------------------------------- //

func (req LoanRequest) isValidAmount() bool {
	return req.Amount > 10000.0
}

func (req LoanRequest) isValidPackage() bool {
	intValue, _ := strconv.Atoi(req.Package)
	return isDigit(req.Package) && intValue >= 500
}

func (req LoanRequest) isValidDuration() bool {
	intValue, _ := strconv.Atoi(req.Duration)
	return isDigit(req.Duration) && intValue >= 6
}

func (req LoanRequest) Validate() *ericerrors.EricError {

	if !req.isValidAmount() {
		return ericerrors.New422Error("Invalid Loan Amount")
	}
	if !req.isValidPackage() {
		return ericerrors.New422Error("Invalid Laon Package")
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
