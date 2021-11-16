package repositories

import (
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/jmoiron/sqlx"
)

type FundsRepo struct {
	db *sqlx.DB
}

// Getter method to inject and get DB instance
func NewFundsRepo(dbClient *sqlx.DB, env utils.Config) FundsRepo {
	return FundsRepo{dbClient}
}

func (repo FundsRepo) FundWallet(payment models.Payment) responsedto.PaymentInitRespnse {
	response := payment.ConvertToFlutterResponseDTO()
	return responsedto.PaymentInitRespnse{
		PaymentBody: response,
		Token:       "",
	}
}

//

func (repo FundsRepo) CompletWalletFunding(funding models.CompleteFunding) (*responsedto.WalletDTO, *ericerrors.EricError) {

	// Before Updated wallet, Check if transaction hasn't been redeemed
	redeemed := hasBeenRedeemed(funding.TxRef, repo.db)
	if redeemed {
		logger.Error("Fraudulent transaction")
		return nil, ericerrors.NewError(400, "Fraudulent transaction")
	} else {
		userWallet := getWallet(funding.Wallet, repo.db)
		prevAmount := userWallet.Amount
		value, Cerr := strconv.ParseFloat(funding.Amount, 64)
		if Cerr != nil {
			logger.Error("Couldn't Make float64")
		}
		newAmount := prevAmount + value
		// Prepare statement
		query := "UPDATE wallet SET amount = ? WHERE id = ?"
		// Hit DB
		_, err := repo.db.Exec(query, newAmount, userWallet.Id)
		// HAndle error case
		if err != nil {
			logger.Error(konstants.QUERY_ERR + err.Error())
			return nil, ericerrors.New500Error(err.Error())
		}
		// Record transaction and check if recorded
		newErr := recordTx(funding.TxRef, repo.db)
		if newErr != nil {
			logger.Error("DID NOT SAVE TRANSACTION")
		}
		// Return the appropriate response
		wallet := responsedto.WalletDTO{Amount: newAmount, Address: userWallet.Address}
		return &wallet, nil
	}

}
