package models

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
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

type FundReopositoryPort interface {
	FundWallet(Payment) responsedto.PaymentInitRespnse
}