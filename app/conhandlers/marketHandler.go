package conhandlers

import (
	"net/http"

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
