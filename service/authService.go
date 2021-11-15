package service

import (
	"net/http"

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
	RegisterUser(requestdto.RegisterRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError)
	VerifyAcc(requestdto.VerifyRequest) (*responsedto.LoginResponseDTO, *ericerrors.EricError)
	Login(requestdto.LoginRequest) (*responsedto.OneUserDto, *ericerrors.EricError)
	CompleteLoginProcess(models.Claim) (*responsedto.CompleteUserDTO, *ericerrors.EricError)
	RequestPasswordChange(models.UserEmail) (*responsedto.OneUserDto, *ericerrors.EricError)
	ChangePassword(requestdto.LoginRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError)
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
func (s UserService) RegisterUser(req requestdto.RegisterRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError) {
	// Validate request
	err := req.ValidateRequest()
	if err != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR)
		return nil, err
	}
	userConstruct := requestdto.BuildUser(req)
	newUser, err := s.repo.Create(userConstruct)
	if err != nil {
		return nil, err
	}
	// create a userDTO with token and otp
	userDTO := getDTOWithTokenAndOTP(newUser, "register")
	utils.SendVerificationMail(userDTO)
	plainResponse := responsedto.PlainResponseDTO{Code: http.StatusOK, Message: konstants.MSG_REG}
	return &plainResponse, nil
}

func (s UserService) Login(req requestdto.LoginRequest) (*responsedto.LoginResponseDTO, *ericerrors.EricError) {
	// Validate request
	err := req.ValidateRequest()
	if err != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR)
		return nil, err
	}

	userLogin := models.UserLogin{Email: req.Email, Password: req.Password}
	dBUser, err := s.repo.Login(userLogin)

	if err != nil {
		logger.Error(konstants.LOGIN_ERR + err.Message)
		return nil, err
	}
	// Check if User pasword matches
	isValidCrentials := security.CheckUserPassword(req.Password, dBUser.Password)
	if !isValidCrentials {
		logger.Error(konstants.LOGIN_ERR)
		return nil, ericerrors.NewCredentialError(konstants.CREDENTIAL_ERR)
	}
	// Construct responseDTOWithToken and responseWithOTP
	userDTO := getDTOWithTokenAndOTP(dBUser, "login")

	// Since all is well, send 2FA otp to user
	utils.SendOTP(userDTO)
	loginResponseDTO := responsedto.LoginResponseDTO{
		Message:     "Check Mail For Login OTP",
		TokenString: userDTO.Token}
	// return appropriate response
	return &loginResponseDTO, nil
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

	res := result.ConvertToVeriyResponse("verified")
	return &res, nil
}

func (s UserService) CompleteLoginProcess(claim models.Claim) (*responsedto.CompleteUserDTO, *ericerrors.EricError) {
	result, err := s.repo.CompleteLogin(claim)
	if err != nil {
		logger.Error(konstants.LOGIN_ERR + err.Message)
		return nil, err
	}
	// Generate Token
	userDTO := getDTOWithTokenAndOTP(result, "complete-login")
	// Convert result to CompleteUser DTO
	return &userDTO, nil
}

func (s UserService) RequestPasswordChange(req requestdto.PasswordChangeRequest) (*responsedto.OneUserDto, *ericerrors.EricError) {
	// Validate request
	err := req.ValidatePwordRequest()
	if err != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR)
		return nil, err
	}
	emailStruct := models.UserEmail{Email: req.Email}
	dBUser, err := s.repo.RequestPasswordChange(emailStruct)

	if err != nil {
		logger.Error(konstants.LOGIN_ERR + err.Message)
		return nil, err
	}

	// Return Queried User with token that will be used for verication on confirm password change
	userDTO := getDTOWithTokenAndOTP(dBUser, "request")
	oneUser := userDTO.GetOneUserFromComplete()
	// Send Mail to user
	utils.SendRequestMail(userDTO)
	return &oneUser, nil

}

// Call Changepassword when user password is to change
func (s UserService) ChangePassword(userReq requestdto.LoginRequest) (*responsedto.PlainResponseDTO, *ericerrors.EricError) {
	err := userReq.ValidateRequest()
	if err != nil {
		logger.Error(konstants.REQ_VALIDITY_ERR)
		return nil, err
	}
	userLogin := models.UserLogin{Email: userReq.Email, Password: userReq.Password}
	res, err := s.repo.ChangePassword(userLogin)
	if err != nil {
		logger.Error("Pword Change Err: " + err.Message)
		return nil, err
	}
	return res, nil
}
