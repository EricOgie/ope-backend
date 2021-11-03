package models

import "time"

type User struct {
	Id        string    `db:"id"`
	FirstName string    `json:"firstname" validate:"required,min=2,max=50" xml:"first_name"`
	LastName  string    `json:"lastname" validate:"required,min=2,max=50" xml:"last_name"`
	Email     string    `json:"email" validate:"email,required" xml:"email"`
	Password  string    `json:"password" xml:"password" validate:"required,min=6"`
	Phone     string    `json:"phone" validate:"required" xml:"phone"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
}

// Add User adapter port
type UserRepositoryPort interface {
	FindAll() ([]User, error)
}
