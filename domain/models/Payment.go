package models

import responsedto "github.com/EricOgie/ope-be/dto/responseDto"

type Payment struct {
	Tx_Ref         string         `json:"tx_ref"`
	Amount         string         `json:"amount"`
	Currency       string         `json:"currency"`
	PaymentOptions string         `json:"payment_options"`
	RedirectUrl    string         `json:"redirect_url"`
	Meta           Meta           `json:"meta"`
	Customer       Customer       `json:"customer"`
	Customizations Customizations `json:"customizations"`
}

type Meta struct {
	ConsumerId  int    `json:"consumer_id"`
	ConsumerMac string `json:"consumer_mac"`
}

type Customer struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Name        string `json:"name"`
}

type Customizations struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

func (p Payment) ConvertToFlutterResponseDTO() responsedto.FlutterResponseDTO {
	return responsedto.FlutterResponseDTO{
		Tx_Ref:         p.Tx_Ref,
		Amount:         p.Amount,
		Currency:       p.Currency,
		PaymentOption:  p.PaymentOptions,
		RedirectUrl:    p.RedirectUrl,
		Meta:           responsedto.MetaDTO(p.Meta),
		Customer:       responsedto.CustomerDTO(p.Customer),
		Customizations: responsedto.CustomizationsDTO(p.Customizations),
	}
}
