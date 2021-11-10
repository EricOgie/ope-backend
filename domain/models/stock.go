package models

type Stock struct {
	Id            string `db:"id" json:"id"`
	Symbol        string `db:"symbol" json:"symbol"`
	ImageUrl      string `db:"image" json:"image_url"`
	QUantity      string `db:"total_quantity" json:"quantity"`
	UnitPrice     string `db:"unit_price" json:"unit_price"`
	Equity        string `db:"equity_value" json:"equity_value"`
	PercentChange string `db:"fluctuation" json:"percentage_change"`
}

func MakeCompleteUser(claim Claim, stocks []Stock) CompleteUser {
	return CompleteUser{
		Id:        claim.Id,
		FirstName: claim.Firstname,
		LastName:  claim.Lastname,
		Email:     claim.Email,
		CreatedAt: claim.When,
		Portfolio: stocks,
	}
}
