package conhandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	response "github.com/EricOgie/ope-be/responses"
)

func makeVerifyReqDTO(claim models.Claim) requestdto.VerifyRequest {
	return requestdto.VerifyRequest{
		Id:         claim.Id,
		FirstName:  claim.Firstname,
		Lastname:   claim.Lastname,
		Email:      claim.Email,
		Created_at: claim.CreatedAt,
		When:       claim.CreatedAt,
	}
}

func IsOTPTheSame(req *http.Request, claim models.Claim) bool {
	var reqOTP requestdto.OTPDto
	err := json.NewDecoder(req.Body).Decode(&reqOTP)

	logger.Info("Claim/Req = " + strconv.Itoa(claim.Otp) + "/" + strconv.Itoa(reqOTP.OTP))

	if err != nil {
		logger.Error(konstants.ERR + err.Error())
	}

	return reqOTP.OTP == claim.Otp

}

func IsOtpValid(otp int) bool {
	stV := strconv.Itoa(otp)
	return len(stV) == 6
}

type PWord struct {
	Password string
}

func handleBadRequest(res http.ResponseWriter) {
	eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
	response.ServeResponse(konstants.ERR, "", res, eError)
}
