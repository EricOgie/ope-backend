package repositories

import (
	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

// Create User Adapter
type UserRepositoryStub struct {
	users *[]responsedto.UserDto
}

// Plug UserRepositoryStub adapter to UserRepository port
func (stub UserRepositoryStub) FindAll() (*[]responsedto.UserDto, error) {
	return stub.users, nil
}

func (stub UserRepositoryStub) Create(models.User) (*models.User, *ericerrors.EricError) {
	return &models.User{
		Id: "1", FirstName: "Todo", LastName: "Didi",
		Email: "jay@gmail.com", Phone: "77336654544", CreatedAt: "2021-02-02",
	}, nil
}

// Helper Function to instantiate UserRepotoryStub
// The function will create and return an instance of UserRepositoryStub
func NewUserRepoStub() UserRepositoryStub {
	// we will create a slice of dummy users
	users := []responsedto.UserDto{
		{Id: "1", FirstName: "Todo", LastName: "Didi", Email: "jay@gmail.com", Phone: "77336654544"},
		{Id: "2", FirstName: "Mesh", LastName: "Jay", Email: "jay@gmail.com", Phone: "77336654544"},
		// {Id: "3", FirstName: "Caleb", LastName: "Ben", Email: "jay@gmail.com", Phone: "77336654544"},
	}
	return UserRepositoryStub{&users}
}
