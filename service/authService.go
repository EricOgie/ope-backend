package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	requestdto "github.com/EricOgie/ope-be/dto/requestDTO"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
	"github.com/EricOgie/ope-be/utils"
)

// Create client side port for User related resource
type UserServicePort interface {
	GetAllUsers() (*[]responsedto.UserDto, error)
	RegisterUser(requestdto.RegisterRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
	Login(requestdto.LoginRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
	VerifyAcc(requestdto.VerifyRequest) (*responsedto.VerifiedRESPONSE, *ericerrors.EricError)
}

// Define UserService as biz end of User domain
type UserService struct {
	repo models.UserRepositoryPort
}

// Helper function to instatiate UserService
//The fnction will create and return an instance of Uservice
func NewUserService(repo models.UserRepositoryPort) UserService {
	return UserService{repo}
}

// Plug userService to UserServicePort via RegisterUser interface implementation
func (s UserService) RegisterUser(req requestdto.RegisterRequest) (*responsedto.OneUserDto, *ericerrors.EricError) {
	// Validate request
	err := req.ValidateRequest()
	if err != nil {
		// hhh

		//jjj
		return nil, err
	}
	userConstruct := requestdto.BuildUser(req)
	newUser, err := s.repo.Create(userConstruct)
	if err != nil {
		return nil, err
	}
	// Add signed token to user struct and return
	userResponseDTOWithToken, resDTOWithToken := getUserWithToken(newUser)
	utils.SendVerificationMail(resDTOWithToken, userResponseDTOWithToken.Token)
	return &userResponseDTOWithToken, nil
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
		logger.Error(konstants.LOGIN_ERR + err.Message)
		return nil, err
	}
	// Check if User pasword matches
	logger.Info("SEE= " + req.Password + "/" + dBUser.Password)
	isValidCrentials := security.CheckUserPassword(req.Password, dBUser.Password)
	if !isValidCrentials {
		logger.Error(konstants.LOGIN_ERR)
		return nil, ericerrors.NewCredentialError(konstants.CREDENTIAL_ERR)
	}
	// Construct responseDTOWithToken and responseWithOTP
	resDTOWithToken, userDTOWithOTP := getUserWithToken(dBUser)
	// Since all is well, send 2FA otp to user
	utils.SendOTP(userDTOWithOTP)

	return &resDTOWithToken, nil
}

// Plug userService to UserServicePort
func (s UserService) GetAllUsers() (*[]responsedto.UserDto, *ericerrors.EricError) {
	return s.repo.FindAll()
}

func (s UserService) VerifyAcc(vr requestdto.VerifyRequest) (*responsedto.VerifiedRESPONSE, *ericerrors.EricError) {
	verifyUser := models.VerifyUser{
		Id:        vr.Id,
		FirstName: vr.FirstName,
		LastName:  vr.Lastname,
		Email:     vr.Email,
		CreatedAt: vr.Created_at,
	}

	result, err := s.repo.VerifyUserAccount(verifyUser)
	// Handle error
	if err != nil {
		logger.Error(konstants.VET_ACC_ERR + err.Message)
		return nil, err
	}

	res := result.ConvertToVeriyResponse()
	return &res, nil
}

//   ----------------------- PRIVATE METHOD ---------------------------- //
func getUserWithToken(user *models.User) (responsedto.OneUserDto, responsedto.OneUserDtoWithOtp) {
	// Gen OTP
	otp := security.GenerateOTP()
	// Construct UserDTOwithOTP from user
	userDTOWithOTP := user.ConvertToOneUserDtoWithOtp(otp)
	// Gen Token
	token := security.GenerateToken(userDTOWithOTP)
	// Contruct UserDTOwithToken
	userResponseDTOWithToken := user.ConvertToOneUserDto(token) //user.ConvertToOneUserDto(token)

	return userResponseDTOWithToken, userDTOWithOTP
}
