package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/service"
)

type UserHandler struct {
	Service service.UserService
}

func (s *UserHandler) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	Users, err := s.Service.GetAllUsers()
	response.ServeResponse("Users Collection", Users, res, err)
}

func (s *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {

	var request requestdto.RegisterRequest
	err := json.NewDecoder(req.Body).Decode(&request)

	// Handle Bad Request Error
	if err != nil {
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: "Bad Request"}
		response.ServeResponse("Error", "me", res, eError)
	} else {
		newUser, eError := s.Service.RegisterUser(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse("User", newUser, res, eError)
	}
}

func (s *UserHandler) Login(res http.ResponseWriter, req *http.Request) {
	var request requestdto.LoginRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	// Handle Bad Request Error
	if err != nil {
		// end process and send 400 error code to client
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: "Bad Request"}
		response.ServeResponse("Error", "", res, eError)
	} else {
		newUser, eError := s.Service.Login(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse("User", newUser, res, eError)
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
		eError := &ericerrors.EricError{Code: http.StatusBadRequest, Message: "Bad Request"}
		response.ServeResponse("Error", "", res, eError)
	}

	response.ServeResponse("User", result, res, nil)
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
