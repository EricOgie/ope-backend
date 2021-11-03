package service

import (
	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/dto"
	"github.com/EricOgie/ope-be/ericerrors"
)

// Create client side port for User related resource
type UserServicePort interface {
	GetAllUsers() (*[]dto.UserDto, error)
}

// Define UserService as biz end of User domain
type UserService struct {
	repo models.UserRepositoryPort
}

// Plug userService to UserServicePort
func (s UserService) GetAllUsers() (*[]dto.UserDto, *ericerrors.EricError) {
	return s.repo.FindAll()
}

// Helper function to instatiate UserService
//The fnction will create and return an instance of Uservice
func NewUserService(repo models.UserRepositoryPort) UserService {
	return UserService{repo}
}
