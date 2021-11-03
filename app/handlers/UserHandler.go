package handlers

import (
	"net/http"

	"github.com/EricOgie/ope-be/helpers"
	"github.com/EricOgie/ope-be/service"
)

type UserHandler struct {
	Service service.UserService
}

func (s *UserHandler) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	Users, err := s.Service.GetAllUsers()
	helpers.ServeResponse(Users, res, err)
}
