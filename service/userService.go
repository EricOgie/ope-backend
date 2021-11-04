package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

// Create client side port for User related resource
type UserServicePort interface {
	GetAllUsers() (*[]responsedto.UserDto, error)
	RegisterUser(requestdto.RegisterRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
}

// Define UserService as biz end of User domain
type UserService struct {
	repo models.UserRepositoryPort
}

// Plug userService to UserServicePort
func (s UserService) GetAllUsers() (*[]responsedto.UserDto, *ericerrors.EricError) {
	return s.repo.FindAll()
}

// Plug userService to UserServicePort via RegisterUser interface implementation
func (s UserService) RegisterUser(req requestdto.RegisterRequest) (*responsedto.OneUserDto, *ericerrors.EricError) {
	// Validate request
	err := req.ValidateRequest()
	if err != nil {
		return nil, err
	}
	userConstruct := requestdto.BuildUser(req)
	newUser, err := s.repo.Create(userConstruct)
	if err != nil {
		return nil, err
	}

	userResponseDTO := newUser.ConvertToOneUserDto("eb267shmdh8.9283hdg76.8769")
	return &userResponseDTO, nil
}

// Helper function to instatiate UserService
//The fnction will create and return an instance of Uservice
func NewUserService(repo models.UserRepositoryPort) UserService {
	return UserService{repo}
}
