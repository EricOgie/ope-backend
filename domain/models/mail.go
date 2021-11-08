package models

type Emailable struct {
	RecipientName  string
	RecipientEmail string
	Subject        string
	Body           string
	IsWithButton   bool
	ButtonText     string
	RedirectUrl    string
	OTP            int
}
