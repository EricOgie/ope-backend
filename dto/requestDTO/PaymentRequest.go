package requestdto

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
)

// To Fund wallet or any other payment, the input from client will be unmashalled into UserPayRequest
// The struct implements a Make make model.Payment methoth that outputs model.Payment neccessary for
// Flutterwave payment
type UserPayRequest struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	PaymentOptions string `json:"payment_option"`
}

type CompleteWalletRequest struct {
	TxRef  string `json:"tx_ref"`
	Wallet string `json:"wallet"`
	Amount string `json:"amount"`
}

type BankRequest struct {
	UserId        string `json:"user_id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_no"`
}

// MakePaymenr is a method implementation on UserPayRequest that converts UserPayRequest to models.Payment struct
// The method takes in an instance of models.Claim struct
func (userPay UserPayRequest) MakePayment(claim models.Claim) models.Payment {
	var wal models.Wallet
	walAsJson, _ := json.Marshal(claim.Wallet)
	_ = json.Unmarshal(walAsJson, &wal)
	userId, _ := strconv.Atoi(claim.Id)
	return models.Payment{
		Tx_Ref:         claim.Firstname + "-tx-" + gencode(),
		Amount:         userPay.Amount,
		Currency:       userPay.Currency,
		PaymentOptions: userPay.PaymentOptions,
		RedirectUrl:    "",
		Meta:           models.Meta{ConsumerId: userId, ConsumerMac: wal.Address},
		Customer:       models.Customer{Email: claim.Email, PhoneNumber: "", Name: claim.Firstname},
		Customizations: models.Customizations{Title: "Fund Wallet", Description: "Funding wallet for subsequent trasaction", Logo: "www.mylogo.com"},
	}
}

func (c CompleteWalletRequest) ConvertToCompletFunding() models.CompleteFunding {
	return models.CompleteFunding{
		TxRef:  c.TxRef,
		Wallet: c.Wallet,
		Amount: c.Amount,
	}
}

// Convert BankRequest to Models.BankAccount struct
func (b BankRequest) ConvertToBankAccount() models.BankAccount {
	return models.BankAccount{UserId: b.UserId, AccountNumber: b.AccountNumber, AccountName: b.BankName}
}

//
func (userPayReq UserPayRequest) IsValidateRequest() bool {
	return userPayReq.IsAmountIsUpto5000() && userPayReq.IsCardOption() && userPayReq.IsNaira()
}

func (req BankRequest) ValidateBankRequest() *ericerrors.EricError {

	if !req.isValidId() {
		return ericerrors.New422Error("Invalid user Id")
	}

	if !req.isValidAccountNumber() {
		return ericerrors.New422Error("Invalid Account Number")
	}

	if !req.isValidBankName() {
		return ericerrors.New422Error("Invalid Bank Name")
	}

	return nil
}

func gencode() string {
	gen := ""
	for i := 0; i < 6; i++ {
		opeRand, _ := rand.Int(rand.Reader, big.NewInt(9))
		gen += opeRand.String()
	}

	return gen
}

//
func (userPayReq UserPayRequest) IsNaira() bool {
	return userPayReq.Currency == "NGN"
}

//
func (userPayReq UserPayRequest) IsAmountIsUpto5000() bool {
	ammount, _ := strconv.Atoi(userPayReq.Amount)
	return ammount >= 1000
}

//
func (userPayReq UserPayRequest) IsCardOption() bool {
	return userPayReq.PaymentOptions == "card"
}

//
func (req CompleteWalletRequest) IsValidTxRef(claim models.PaymentClaim) bool {
	return len(req.TxRef) >= 11 && req.TxRef == claim.TxRef
}

func (req CompleteWalletRequest) IsValidWallet() bool {
	return len(req.Wallet) > 40
}

func (req CompleteWalletRequest) IsValidAmount(claim models.PaymentClaim) bool {
	return req.Amount == claim.Amount
}

func isDigit(value string) bool {
	_, e := strconv.Atoi(value)
	return e == nil
}

func (r BankRequest) isValidId() bool {
	return r.UserId != "0" && len(r.UserId) > 0
}

func (r BankRequest) isValidAccountNumber() bool {
	return isDigit(r.AccountNumber) && len(r.AccountNumber) >= 10
}

func (r BankRequest) isValidBankName() bool {
	return len(r.BankName) > 5
}
