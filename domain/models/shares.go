package models

import "github.com/EricOgie/ope-be/ericerrors"

type Share struct {
	Id            string `db:"id" json:"id"`
	Symbol        string `db:"symbol" json:"symbol"`
	ImageUrl      string `db:"image" json:"image_url"`
	UnitPrice     string `db:"unit_price" json:"unit_price"`
	PercentChange string `db:"percentage_change" json:"percentage_change"`
}

type MarketRepositoryPort interface {
	ShowStockMarket() (*[]Share, *ericerrors.EricError)
}
