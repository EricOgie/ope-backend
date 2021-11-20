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
	_, err := s.Service.VerifyAcc(verifyRequest)

	if err != nil {
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	}

	// response.ServeResponse(konstants.USER, result, res, nil)

	response.RedirectToVerified(res, req)
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

func (s *UserHandler) RequestPasswordChange(res http.ResponseWriter, req *http.Request) {
	var request requestdto.PasswordChangeRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	// Handle Bad Request Error
	if err != nil {
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	} else {
		newUser, eError := s.Service.RequestPasswordChange(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse("Request", newUser, res, eError)
	}
}

func (s *UserHandler) ChangePassword(res http.ResponseWriter, req *http.Request) {
	var pword PWord
	err1 := json.NewDecoder(req.Body).Decode(&pword)
	if err1 != nil {
		// User did not enter a valid password
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	}
	// User email will be gotten from claim extracted from token
	claim, _ := req.Context().Value(konstants.DT_KEY).(models.Claim)
	UserCrendentials := requestdto.LoginRequest{Email: claim.Email, Password: pword.Password}
	newUser, eError := s.Service.ChangePassword(UserCrendentials)
	response.ServeResponse("Plain Response", newUser, res, eError)
}

func (s *UserHandler) UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	idAsInt := getUserId(req)
	var request requestdto.UserDetailsRequest
	reqErr := json.NewDecoder(req.Body).Decode(&request)

	if reqErr != nil {
		logger.Error(konstants.ERR_DECODE + reqErr.Error())
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: konstants.BAD_REQ}
		response.ServeResponse(konstants.ERR, "", res, eError)
	}

	// Define the Id attribute of the request
	request.Id = idAsInt
	result, eError := s.Service.ProfileUpdate(request)
	response.ServeResponse(konstants.USER, result, res, eError)

}

func (s *UserHandler) UpdateUserBank(res http.ResponseWriter, req *http.Request) {
	userId := getUserId(req)
	var request requestdto.BankRequest
	jErr := json.NewDecoder(req.Body).Decode(&request)

	if jErr != nil {
		handleMarshallingErr(res, jErr)
	} else {
		request.UserId = strconv.Itoa(userId)
		result, ericErr := s.Service.SetBankDetails(request)
		response.ServeResponse("Account", result, res, ericErr)
	}

}

func (s *UserHandler) FindOneUser(res http.ResponseWriter, req *http.Request) {
	// Extract claim from request and get user email
	claim := req.Context().Value(konstants.DT_KEY).(models.Claim)
	userEmail := claim.Email

	result, err := s.Service.FetchOneUser(userEmail)
	response.ServeResponse(konstants.USER, result, res, err)
}
