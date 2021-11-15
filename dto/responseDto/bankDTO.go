package responsedto

type BankAccountDTO struct {
	AccountNo   string `json:"account_no"`
	AccountName string `json:"bank_name"`
}

type WalletDTO struct {
	Amount  float64 `json:"amount"`
	Address string  `json:"address"`
}
