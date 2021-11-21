package conhandlers

import (
	"encoding/json"
	"net/http"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/service"
)

type MarkHandler struct {
	Service service.MarketService
}

func (handler MarkHandler) FetchMarketState(res http.ResponseWriter, req *http.Request) {
	result, err := handler.Service.ShowStockMarket()
	response.ServeResponse("Shares", result, res, err)
}

func (handler MarkHandler) BuyInvestment(res http.ResponseWriter, req *http.Request) {
	userId := getUserIdAsString(req)
	var request requestdto.PurchaseRequest
	mErr := json.NewDecoder(req.Body).Decode(&request)

	if mErr != nil {
		logger.Error(konstants.ERR + mErr.Error())
		handleBadRequest(res)
	} else {
		request.UserId = userId
		result, err := handler.Service.PurchaseStock(request)
		response.ServeResponse("Plain Response", result, res, err)
	}

}
