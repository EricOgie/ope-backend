package repositories

import (
	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
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
