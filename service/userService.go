package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
)

// Create client side port for User related resource
type UserServicePort interface {
	GetAllUsers() (*[]responsedto.UserDto, error)
	RegisterUser(requestdto.RegisterRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
	Login(requestdto.LoginRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
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
	// Add signed token to user struct and return
	userResponseDTOWithToken := getUserWithToken(newUser)
	return &userResponseDTOWithToken, nil
}

// Helper function to instatiate UserService
//The fnction will create and return an instance of Uservice
func NewUserService(repo models.UserRepositoryPort) UserService {
	return UserService{repo}
}

func (s UserService) Login(req requestdto.LoginRequest) (*responsedto.OneUserDto, *ericerrors.EricError) {
	// Validate request
	err := req.ValidateRequest()
	if err != nil {
		return nil, err
	}

	userLogin := models.UserLogin{Email: req.Email, Password: req.Password}
	dBUser, err := s.repo.Login(userLogin)

	if err != nil {
		return nil, err
	}
	// Check if User pasword matches
	security.CheckUserPassword(req.Password, dBUser.Password)
	logger.Debug("HASH PWORD = " + dBUser.Password)
	// add token to user struct and return
	userResponseDTOWithToken := getUserWithToken(dBUser)
	return &userResponseDTOWithToken, nil
}

//   ----------------------- PRIVATE METHOD ---------------------------- //
func getUserWithToken(user *models.User) responsedto.OneUserDto {
	userResponseDTO := user.ConvertToOneUserDto("")
	token := security.GenerateToken(userResponseDTO)
	userResponseDTOWithToken := user.ConvertToOneUserDto(token)
	return userResponseDTOWithToken
}
