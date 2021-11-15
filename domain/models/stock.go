package models

import (
	"encoding/json"
	"fmt"

	"github.com/EricOgie/ope-be/logger"
)

type Stock struct {
	Id            string `db:"id" json:"id"`
	Symbol        string `db:"symbol" json:"symbol"`
	ImageUrl      string `db:"image" json:"image_url"`
	QUantity      string `db:"total_quantity" json:"quantity"`
	UnitPrice     string `db:"unit_price" json:"unit_price"`
	Equity        string `db:"equity_value" json:"equity_value"`
	PercentChange string `db:"percentage_change" json:"percentage_change"`
}

func MakeCompleteUser(claim Claim, stocks []Stock) CompleteUser {
	logger.Info("JUST TO SHOW BEFORE")
	fmt.Println(fmt.Sprintf("%#v", claim.BankAccount))
	var acc BankAccount
	var wal Wallet
	valInJson, _ := json.Marshal(claim.BankAccount)
	walJson, _ := json.Marshal(claim.Wallet)

	_ = json.Unmarshal(valInJson, &acc)
	_ = json.Unmarshal(walJson, &wal)

	return CompleteUser{
		Id:          claim.Id,
		FirstName:   claim.Firstname,
		LastName:    claim.Lastname,
		Email:       claim.Email,
		CreatedAt:   claim.CreatedAt,
		BankAccount: acc,
		Wallet:      wal,
		Portfolio:   stocks,
	}
}
