package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
)

type MarketServicePort interface {
	ShowStockMarket() (*[]models.Share, *ericerrors.EricError)
}

type MarketService struct {
	Repo models.MarketRepositoryPort
}

// Finds market stocks in their current valuation from db
func (service MarketService) ShowStockMarket() (*[]models.Share, *ericerrors.EricError) {
	return service.Repo.ShowStockMarket()
}

// Helper function to Get and instance of marketService
func NewMarketService(repo models.MarketRepositoryPort) MarketService {
	return MarketService{repo}
}
