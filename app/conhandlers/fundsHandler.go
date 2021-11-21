package conhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/service"
)

type FundHandler struct {
	Service service.PaymentService
}

func (handler FundHandler) FundUserWallet(res http.ResponseWriter, req *http.Request) {
	// extract claim from http.Request
	claim, _ := req.Context().Value(konstants.DT_KEY).(models.Claim)

	var request requestdto.UserPayRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		handleBadRequest(res)
	} else {
		// call service
		result, rErr := handler.Service.FundWallet(request, claim)
		// Pass response to response layer to serve the appropriate response
		response.ServeResponse("FultterWave", result, res, rErr)
	}

}

//
func (handler FundHandler) CompleteFundingFlow(res http.ResponseWriter, req *http.Request) {

	// Retreve claim and cast to payment claim structS
	claim, _ := req.Context().Value(konstants.DT_KEY).(models.PaymentClaim)
	var request requestdto.CompleteWalletRequest
	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		handleBadRequest(res)
	} else {
		// vallidate request
		isvalidRe := request.IsValidAmount(claim) && request.IsValidAmount(claim) && request.IsValidTxRef(claim)
		if !isvalidRe {
			handleBadRequest(res)
			return
		}

		result, eErr := handler.Service.CompleteFunding(request)
		if eErr != nil {
			response.ServeResponse(konstants.ERR, "", res, eErr)
		}

		response.ServeResponse("Wallet", result, res, eErr)
	}

}
