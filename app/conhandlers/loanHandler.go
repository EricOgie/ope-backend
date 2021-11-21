package conhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/service"
)

type LoanHandler struct {
	Service service.LoanService
}

//
//
func (h LoanHandler) RequestLoan(res http.ResponseWriter, req *http.Request) {
	id := getUserIdFromClaim(req)
	var request requestdto.LoanRequest

	mErr := json.NewDecoder(req.Body).Decode(&request)
	if mErr != nil {
		logger.Error(konstants.ERR + mErr.Error())
		handleBadRequest(res)
	} else {
		request.UserId = id
		result, err := h.Service.Borrow(request)
		response.ServeResponse("Loan", result, res, err)
	}
}

//
//
func (h LoanHandler) GetAllUserLoans(res http.ResponseWriter, req *http.Request) {
	id := convertStringToInt(getUserIdFromClaim(req))
	result, err := h.Service.GetLoans(id)
	response.ServeResponse("Loan Collection", result, res, err)
}

//
//
func (h LoanHandler) RePayInInstallment(res http.ResponseWriter, req *http.Request) {
	loanId := getLoanIdAsString(req)
	claim := req.Context().Value(konstants.DT_KEY).(models.Claim)
	var request requestdto.LoanPayRequest

	mErr := json.NewDecoder(req.Body).Decode(&request)
	if mErr != nil {
		logger.Error(konstants.ERR + mErr.Error())
		handleBadRequest(res)
	} else {
		request.UserId = claim.Id
		request.LoanId = loanId

		result, err := h.Service.PayInPart(request)
		response.ServeResponse("Repayment", result, res, err)
	}
}

//
//
func (h LoanHandler) GetLoanPayments(res http.ResponseWriter, req *http.Request) {
	loanId := getLoanId(req)
	result, err := h.Service.FetchInstallments(loanId)
	response.ServeResponse("Loan-Payments", result, res, err)
}
