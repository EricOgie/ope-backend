package handlers

import (
	"encoding/json"
	"net/http"

	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/logger"
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
		// response.ServeError(http.StatusBadRequest, err.Error(), res)
	} else {
		logger.Debug("H1")
		newUser, eError := s.Service.RegisterUser(request)
		// Send response and Error to Response handler layer and allow
		//it serve the appropriate response to client
		response.ServeResponse("User", newUser, res, eError)
	}
}
