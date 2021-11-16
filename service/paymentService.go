package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
)

type PaymentServicePort interface {
	FundWallet(requestdto.UserPayRequest, models.Claim) *responsedto.FlutterResponseDTO
}
type PaymentService struct {
	repo models.FundReopositoryPort
}

// Getter func dto instantiate ad to get PaymentService
func NewPaymentService(payPort models.FundReopositoryPort) PaymentService {
	return PaymentService{payPort}
}

// FundWallet is a payment method that runs to FundsRepo for implementation
// It takes requestdto.UserPayRequest, and models.Claim and spill result directly from FundsRepo actions
func (payService PaymentService) FundWallet(userPayReq requestdto.UserPayRequest,
	claim models.Claim) (*responsedto.PaymentInitRespnse, *ericerrors.EricError) {

	if !userPayReq.IsValidateRequest() {
		logger.Error(konstants.REQ_VALIDITY_ERR)
		return nil, ericerrors.New422Error(konstants.PAY_VALIDATION_ERR_MSG)
	}
	// Convert UserPayRequest to models.Payment
	payment := userPayReq.MakePayment(claim)
	// call Repo
	result := payService.repo.FundWallet(payment)
	result.Token = security.GenPaymentToken(&result.PaymentBody)
	// Return correct response
	return &result, nil
}
