package models

import "github.com/EricOgie/ope-be/ericerrors"

type Share struct {
	Id            string  `db:"id" json:"id"`
	Symbol        string  `db:"symbol" json:"symbol"`
	ImageUrl      string  `db:"image" json:"image_url"`
	UnitPrice     float64 `db:"unit_price" json:"unit_price"`
	PercentChange float64 `db:"percentage_change" json:"percentage_change"`
	Volume        float64 `db:"volume" json:"volume"`
}

type MarketRepositoryPort interface {
	ShowStockMarket() (*[]Share, *ericerrors.EricError)
}
