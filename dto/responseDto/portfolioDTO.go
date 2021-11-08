package responsedto

type StockDTO struct {
	Id            string `json:"stock_id"`
	Symbol        string `json:"symbol"`
	ImageUrl      string `json:"image"`
	QUantity      string `json:"quantity"`
	UnitPrice     string `json:"unit_price"`
	Equity        string `json:"equity"`
	PercentChange string `json:"percent_change"`
}
