package models

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

type Wallet struct {
	Id      int64   `db:"id" json:"id"`
	Amount  float64 `db:"amount" json:"amount"`
	Address string  `db:"address" json:"address"`
}

type Fund struct {
	Amount float64 `db:"amount" json:"amount"`
	UserId int64   `db:"user_id" json:"address"`
}

type CompleteFunding struct {
	TxRef  string
	Wallet string
	Amount string
}

type FundReopositoryPort interface {
	FundWallet(Payment) responsedto.PaymentInitRespnse
	CompletWalletFunding(funding CompleteFunding) (*responsedto.WalletDTO, *ericerrors.EricError)
}
