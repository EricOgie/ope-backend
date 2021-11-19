package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
)

type LoanServicePort interface {
	Borrow(requestdto.LoanRequest) (*responsedto.LoanResDTO, *ericerrors.EricError)
	GetLoans(int) (*[]models.QueryLoan, *ericerrors.EricError)
	PayInPart(requestdto.LoanPayRequest) (*responsedto.RepaymentResDTO, *ericerrors.EricError)
	FetchInstallments(int) (*[]models.Querypayment, *ericerrors.EricError)
}

type LoanService struct {
	repo models.LoanRepositoryPort
}

func NewLoanService(repo models.LoanRepositoryPort) LoanService {
	return LoanService{repo}
}

//
func (s LoanService) Borrow(req requestdto.LoanRequest) (*responsedto.LoanResDTO, *ericerrors.EricError) {

	vErr := req.Validate()
	if vErr != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR + vErr.Message)
		return nil, vErr
	}

	loanObj := req.ConvertToLoan()
	result, err := s.repo.TakeLoan(loanObj)

	if err != nil {
		return nil, err
	}

	return result, nil

}

//
func (s LoanService) GetLoans(userId int) (*[]models.QueryLoan, *ericerrors.EricError) {
	result, err := s.repo.FetchLoans(userId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

//
func (s LoanService) PayInPart(req requestdto.LoanPayRequest) (*responsedto.RepaymentResDTO, *ericerrors.EricError) {
	vErr := req.Validate()
	if vErr != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR + vErr.Message)
		return nil, vErr
	}

	loanReq := req.ConvertToLoanPayment()
	result, err := s.repo.PayLoan(loanReq)

	if err != nil {
		return nil, err
	}

	return result, nil

}

//
func (s LoanService) FetchInstallments(loanId int) (*[]models.Querypayment, *ericerrors.EricError) {
	result, err := s.repo.GetInstallments(loanId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
