package models

import (
	"strconv"

	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

type Loan struct {
	Id        int     `db:"id" json:"id"`
	UserId    int     `db:"user_id" json:"user_id"`
	Amount    float64 `db:"amount" json:"amount"`
	Paid      float64 `db:"paid" json:"paid"`
	Package   string  `db:"package" json:"package"`
	Duration  string  `db:"duration" json:"duration"`
	Status    string  `db:"status" json:"status"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}

//
type Repayment struct {
	Id       int     `db:"id" json:"id"`
	UserId   int     `db:"user_id" json:"user_id"`
	Payment  float64 `db:"payment" json:"payment"`
	Amount   float64 `db:"amount" json:"amount"`
	Package  string  `db:"package" json:"package"`
	Duration string  `db:"duration" json:"duration"`
	Status   string  `db:"status" json:"status"`
	Date     string  `db:"created_at" json:"date"`
}

//
type Querypayment struct {
	Id      int     `db:"id" json:"id"`
	Payment float64 `db:"payment" json:"payment"`
	Balance float64 `db:"balance" json:"balance"`
	Amount  float64 `db:"amount" json:"amount"`
	Date    string  `db:"created_at" json:"date"`
}

type LoanPayment struct {
	LoanId    int     `json:"loan_id"`
	UserId    int     `json:"user_id"`
	Payment   float64 `json:"payment"`
	CreatedAt string  `json:"created_at"`
}

type LoanRepositoryPort interface {
	TakeLoan(Loan) (*responsedto.LoanResDTO, *ericerrors.EricError)
	FetchLoans(int) (*[]Loan, *ericerrors.EricError)
	PayLoan(LoanPayment) (*responsedto.RepaymentResDTO, *ericerrors.EricError)
	GetInstallments() (*[]Querypayment, *ericerrors.EricError)
}

//
func (l Loan) ConvertToLoanResDTO() responsedto.LoanResDTO {
	return responsedto.LoanResDTO{
		Id:       strconv.Itoa(l.Id),
		Amount:   l.Amount,
		Paid:     l.Paid,
		Package:  l.Package,
		Duration: l.Duration,
		Status:   l.Status,
		Date:     l.CreatedAt,
	}
}

//
func (r Repayment) ConvertToRepaymentDTO(paid float64) responsedto.RepaymentResDTO {
	return responsedto.RepaymentResDTO{
		Id:      strconv.Itoa(r.Id),
		Payment: r.Payment,
		Balance: (r.Amount - (r.Payment + paid)),
		Amount:  r.Amount,
		Date:    r.Date,
		Status:  r.Status,
	}
}

//
func (r LoanPayment) ConvertToRepaymentDTO(paid float64, installmentId int, loan *Loan, status string) responsedto.RepaymentResDTO {
	return responsedto.RepaymentResDTO{
		Id:      strconv.Itoa(installmentId),
		Payment: paid,
		Balance: (loan.Amount - (loan.Paid + paid)),
		Amount:  loan.Amount,
		Date:    r.CreatedAt,
		Status:  status,
	}
}
