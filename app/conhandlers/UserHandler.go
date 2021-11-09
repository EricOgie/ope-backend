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
	"github.com/EricOgie/ope-be/service"
)

type UserHandler struct {
	Service service.UserService
}

func (s *UserHandler) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	Users, err := s.Service.GetAllUsers()
	response.ServeResponse(konstants.USER_COLL, Users, res, err)
}

func (s *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var request requestdto.RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&request)

	// Handle Bad Request Error
	if err != nil {
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	} else {
		newUser, eError := s.Service.RegisterUser(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse(konstants.USER, newUser, res, eError)
	}
}

func (s *UserHandler) Login(res http.ResponseWriter, req *http.Request) {
	var request requestdto.LoginRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	// Handle Bad Request Error
	if err != nil {
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	} else {
		newUser, eError := s.Service.Login(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse(konstants.LOGIN, newUser, res, eError)
	}
}

func (s *UserHandler) VerifyUserAcc(res http.ResponseWriter, req *http.Request) {
	// access the intent claim from the request
	claim, _ := req.Context().Value(konstants.DT_KEY).(models.Claim)
	// construct a verifyRequest from models.Claim
	verifyRequest := makeVerifyReqDTO(claim)
	//make request along the wiring chain
	result, err := s.Service.VerifyAcc(verifyRequest)

	if err != nil {
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	}

	response.ServeResponse(konstants.USER, result, res, nil)
}

func (s *UserHandler) CompleteLoginProcess(res http.ResponseWriter, req *http.Request) {

	// access the intent claim from the request
	claim, _ := req.Context().Value(konstants.DT_KEY).(models.Claim)

	if !IsOTPTheSame(req, claim) {
		response.ServeResponse(konstants.ERR, "", res, &ericerrors.EricError{
			Code:    http.StatusForbidden,
			Message: konstants.FORBIDDEN,
		})
	} else {
		// Initiate process
		result, err := s.Service.CompleteLoginProcess(claim)
		if err != nil {
			logger.Error(konstants.LOGIN_ERR + err.Message)
			response.ServeResponse(konstants.ERR, "", res, err)
		}

		// send ok response
		response.ServeResponse(konstants.USER, result, res, err)
	}

}

func makeVerifyReqDTO(claim models.Claim) requestdto.VerifyRequest {
	return requestdto.VerifyRequest{
		Id:         claim.Id,
		FirstName:  claim.Firstname,
		Lastname:   claim.Lastname,
		Email:      claim.Email,
		Created_at: claim.When,
		When:       claim.When,
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
