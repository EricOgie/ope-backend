package repositories

import (
	"github.com/EricOgie/ope-be/domain/models"
)

// Create User Adapter
type UserRepositoryStub struct {
	users *[]models.User
}

// Plug UserRepositoryStub adapter to UserRepository port
func (stub UserRepositoryStub) FindAll() (*[]models.User, error) {
	return stub.users, nil
}

// Helper Function to instantiate UserRepotoryStub
// The function will create and return an instance of UserRepositoryStub
func NewUserRepoStub() UserRepositoryStub {
	// we will create a slice of dummy users
	users := []models.User{
		{Id: "1", FirstName: "Todo", LastName: "Didi", Email: "jay@gmail.com", Password: "12345678", Phone: "77336654544"},
		{Id: "2", FirstName: "Mesh", LastName: "Jay", Email: "jay@gmail.com", Password: "12345678", Phone: "77336654544"},
		// {Id: "3", FirstName: "Caleb", LastName: "Ben", Email: "jay@gmail.com", Phone: "77336654544"},
	}
	return UserRepositoryStub{&users}
}
