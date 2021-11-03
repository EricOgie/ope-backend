package repositories

import (
	"github.com/EricOgie/ope-be/dto"
)

// Create User Adapter
type UserRepositoryStub struct {
	users *[]dto.UserDto
}

// Plug UserRepositoryStub adapter to UserRepository port
func (stub UserRepositoryStub) FindAll() (*[]dto.UserDto, error) {
	return stub.users, nil
}

// Helper Function to instantiate UserRepotoryStub
// The function will create and return an instance of UserRepositoryStub
func NewUserRepoStub() UserRepositoryStub {
	// we will create a slice of dummy users
	users := []dto.UserDto{
		{Id: "1", FirstName: "Todo", LastName: "Didi", Email: "jay@gmail.com", Phone: "77336654544"},
		{Id: "2", FirstName: "Mesh", LastName: "Jay", Email: "jay@gmail.com", Phone: "77336654544"},
		// {Id: "3", FirstName: "Caleb", LastName: "Ben", Email: "jay@gmail.com", Phone: "77336654544"},
	}
	return UserRepositoryStub{&users}
}
