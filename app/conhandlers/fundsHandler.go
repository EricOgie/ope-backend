package conhandlers

import (
	"encoding/json"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/ericerrors"
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
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	}

	// call service
	result, rErr := handler.Service.FundWallet(request, claim)

	// Pass response to response layer to serve the appropriate response
	response.ServeResponse("FultterWave Response", result, res, rErr)
}
