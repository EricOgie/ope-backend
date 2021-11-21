package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
)

type MarketServicePort interface {
	ShowStockMarket() (*[]models.Share, *ericerrors.EricError)
	PurchaseStock(requestdto.PurchaseRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError)
}

type MarketService struct {
	Repo models.MarketRepositoryPort
}

// Helper function to Get and instance of marketService
func NewMarketService(repo models.MarketRepositoryPort) MarketService {
	return MarketService{repo}
}

// Finds market stocks in their current valuation from db
func (service MarketService) ShowStockMarket() (*[]models.Share, *ericerrors.EricError) {
	return service.Repo.ShowStockMarket()
}

//
func (service MarketService) PurchaseStock(req requestdto.PurchaseRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError) {
	// Validate
	vErr := req.ValidateRequest()
	if vErr != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR + vErr.Message)
		return nil, vErr
	}

	buystock := req.ConvertToShareStock()
	result, err := service.Repo.BuyStock(buystock)
	// REturn appropriate response
	if err != nil {
		return nil, err
	}

	return result, nil
}
