package models

import (
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
)

type User struct {
	Id        string `db:"id"`
	FirstName string `json:"firstname" validate:"required,min=2,max=50" xml:"first_name"`
	LastName  string `json:"lastname" validate:"required,min=2,max=50" xml:"last_name"`
	Email     string `json:"email" validate:"email,required" xml:"email"`
	Phone     string `json:"phone" validate:"required" xml:"phone"`
	Password  string `json:"password" xml:"password" validate:"required,min=6"`
	CreatedAt string `json:"created_at" validate:"required" xml:"created_at"`
	UpdatedAt string `json:"update_at" validate:"required" xml:"updated_at"`
}

// Add User adapter port
type UserRepositoryPort interface {
	FindAll() (*[]responsedto.UserDto, *ericerrors.EricError)
	Create(User) (*User, *ericerrors.EricError)
}

/**
* When serving user data to client side, it would be bad practice to send
* sensitive data like hashed user password alongside. Hence, data access object
* is used here
 */
// Getter function to conver User struct to UserDTO struc
func (user User) ConvertToUserDto() responsedto.UserDto {
	return responsedto.UserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user User) ConvertToOneUserDto(token string) responsedto.OneUserDto {
	return responsedto.OneUserDto{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}

}
