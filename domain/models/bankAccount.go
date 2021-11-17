package models

type BankAccount struct {
	UserId        string `db:"user_id" json:"user_id"`
	AccountNumber string `db:"account_name" json:"account_name"`
	AccountName   string `db:"account_no" json:"account_no"`
}

type QueryUser struct {
	Id          string  `db:"id" json:"id"`
	FirstName   string  `db:"firstname" json:"firstname"`
	LastName    string  `db:"lastname" json:"lastname"`
	Email       string  `db:"email" json:"email"`
	Phone       string  `db:"phone" json:"phone"`
	Password    string  `db:"password" json:"password"`
	CreatedAt   string  `db:"created_at"`
	AccountNo   string  `db:"account_no" json:"account_no"`
	AccountName string  `db:"account_name" json:"account_name"`
	Amount      float64 `db:"amount" json:"amount"`
	Address     string  `db:"address" json:"address"`
}

// MakeAllInOneUserDTO function will output a complete user dTO with account, wallet and portfolio slice
func (qUser QueryUser) MakeCompleteUserFromQueryUser() CompleteUser {
	return CompleteUser{
		Id:          qUser.Id,
		FirstName:   qUser.FirstName,
		LastName:    qUser.LastName,
		Email:       qUser.Email,
		Password:    qUser.Password,
		CreatedAt:   qUser.CreatedAt,
		BankAccount: BankAccount{AccountName: qUser.AccountName, AccountNumber: qUser.AccountNo},
		Wallet:      Wallet{Amount: qUser.Amount, Address: qUser.Address},
	}
}
