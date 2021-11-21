package requestdto

import (
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
)

type PurchaseRequest struct {
	UserId        string
	Symbol        string  `json:"symbol"`
	ImageUrl      string  `json:"image_url"`
	QUantity      string  ` json:"quantity"`
	UnitPrice     float64 ` json:"unit_price"`
	PercentChange float64 `json:"percentage_change"`
}

func (p PurchaseRequest) ConvertToShareStock() models.ShareStock {
	qty, _ := strconv.Atoi(p.QUantity)
	equity := p.UnitPrice * float64(qty)
	return models.ShareStock{
		OwnerId:       p.UserId,
		Symbol:        p.Symbol,
		ImageUrl:      p.ImageUrl,
		QUantity:      p.QUantity,
		UnitPrice:     p.UnitPrice,
		Equity:        equity,
		PercentChange: p.PercentChange,
	}
}

// Validation

func (p PurchaseRequest) ValidateRequest() *ericerrors.EricError {
	if !p.isValidQty() {
		return ericerrors.New422Error("Invalid Quantity")
	}

	if !p.isValidPrice() {
		return ericerrors.New422Error("Invalid Price")
	}

	return nil
}

func (p PurchaseRequest) isValidQty() bool {
	return isDigit(p.QUantity) && p.QUantity != "0"
}

func (p PurchaseRequest) isValidPrice() bool {
	return p.UnitPrice > 0
}
