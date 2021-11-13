package repositories

import (
	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/jmoiron/sqlx"
)

type MarketRepoDB struct {
	client *sqlx.DB
}

// Helper function to get an instance of sqlx DB
func NewMarketRepoDB(dbClient *sqlx.DB, env utils.Config) MarketRepoDB {
	return MarketRepoDB{dbClient}
}

/**
* @ShowStockMarket
* METHOD implemetation of MarketRepositoryPort as an interface
* Interface implementation here correspond to Plugging  UserRepositoryDB
* adapter to MarketRepositoryPort
* Should be callable only when active user is correctly loggedIn
 */
func (db MarketRepoDB) ShowStockMarket() (*[]models.Share, *ericerrors.EricError) {
	marketstock := make([]models.Share, 0)
	query := "SELECT * FROM shares"
	err := db.client.Select(&marketstock, query)

	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	return &marketstock, nil
}
