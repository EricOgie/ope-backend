package repositories

import (
	"net/http"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
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

func (db MarketRepoDB) BuyStock(s models.ShareStock) (*responsedto.PlainResponseDTO, *ericerrors.EricError) {
	ownerId, _ := strconv.Atoi(s.OwnerId)
	quantity, _ := strconv.ParseFloat(s.QUantity, 64)
	amount := s.UnitPrice * quantity

	// Checks funds surficiency
	fundSurficiency := checkSurficientFunds(db, amount, ownerId)
	if !fundSurficiency {
		logger.Error(konstants.ERR_INSURFICIENCY)
		return nil, ericerrors.NewError(http.StatusNotAcceptable, konstants.ERR_INSURFICIENCY)
	}

	// Record stock and handle error cases
	queryErr := recordStock(db, s, ownerId, quantity)
	if queryErr != nil {
		logger.Error(konstants.QUERY_ERR + queryErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// Less amount from wallet balance
	lErr := lessWalletAmount(db, s)
	if lErr != nil {
		logger.Error(" Wallet Update Err: " + lErr.Message)
	}

	return &responsedto.PlainResponseDTO{Code: http.StatusOK, Message: "Succsssfully purchased " + s.Symbol}, nil
}

// ----------------------------------- PRIVATE METHODS --------------------------------------------//
//

func recordStock(db MarketRepoDB, s models.ShareStock, ownerId int, quantity float64) error {

	hasStock := hasStockPrior(db, ownerId, s.Symbol)
	var queryErr error
	// Register stock purchased base on user user transaction history
	if hasStock {

		query := "UPDATE stocks SET total_quantity = total_quantity + ?, equity_value = equity_value + ?" +
			" WHERE user_id = ?"
		_, queryErr = db.client.Exec(query, quantity, s.Equity, ownerId)

	} else {
		query := "INSERT INTO stocks (user_id, symbol, image, total_quantity, unit_price," +
			" equity_value, percentage_change) VALUE (?, ?, ?, ?, ?, ?, ?)"
		_, queryErr = db.client.Exec(query, ownerId, s.Symbol, s.ImageUrl, quantity, s.UnitPrice,
			s.Equity, s.PercentChange)
	}
	return queryErr
}
