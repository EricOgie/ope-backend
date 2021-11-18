package models

import (
	"encoding/json"
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

type ShareStock struct {
	OwnerId       string
	Symbol        string  `json:"symbol"`
	ImageUrl      string  `json:"image_url"`
	QUantity      string  ` json:"quantity"`
	UnitPrice     float64 ` json:"unit_price"`
	Equity        float64 `json:"equity_value"`
	PercentChange float64 `json:"percentage_change"`
}

func MakeCompleteUser(claim Claim, stocks []Stock) CompleteUser {

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
		Holdings:    claim.Holdings,
		BankAccount: acc,
		Wallet:      wal,
		Portfolio:   stocks,
	}
}
