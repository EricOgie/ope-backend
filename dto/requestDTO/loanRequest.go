package requestdto

import (
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
)

type LoanPayRequest struct {
	UserId  string  `json:"user_id"`
	LoadId  string  `json:"loan_id"`
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
	loadId, _ := strconv.Atoi(req.LoadId)

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
		Package:   req.Package,
		Duration:  req.Duration,
		CreatedAt: time.Now().String(),
	}
}
