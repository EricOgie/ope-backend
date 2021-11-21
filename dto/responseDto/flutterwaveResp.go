package responsedto

type MetaDTO struct {
	ConsumerId  int    `json:"consumer_id"`
	ConsumerMac string `json:"consumer_mac"`
}

type CustomizationsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type CustomerDTO struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Name        string `json:"name"`
}

type FlutterResponseDTO struct {
	Tx_Ref         string            `json:"tx_ref"`
	Amount         string            `json:"amount"`
	Currency       string            `json:"currency"`
	PaymentOption  string            `json:"payment_option"`
	RedirectUrl    string            `json:"redirect_url"`
	Meta           MetaDTO           `json:"meta"`
	Customer       CustomerDTO       `json:"customer"`
	Customizations CustomizationsDTO `json:"customizations"`
}

type PaymentInitRespnse struct {
	PaymentBody FlutterResponseDTO `json:"payment_body"`
	Token       string             `json:"token"`
}
