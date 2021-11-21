package responsedto

type LoanResDTO struct {
	Id       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Paid     float64 `json:"paid"`
	Package  string  `json:"package"`
	Duration string  `son:"duration"`
	Status   string  `json:"status"`
	Date     string  `json:"date"`
}

type RepaymentResDTO struct {
	Id      string  `json:"id"`
	Payment float64 `json:"payment"`
	Balance float64 `json:"balance"`
	Amount  float64 `json:"amount"`
	Date    string  `json:"date"`
	Status  string  `json:"status"`
}
