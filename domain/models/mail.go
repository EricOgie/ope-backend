package models

type Emailable struct {
	RecipientName  string
	RecipientEmail string
	Subject        string
	Caption        string
	Body           string
	Tail           string
	IsWithButton   bool
	ButtonText     string
	RedirectUrl    string
	OTP            string
}
