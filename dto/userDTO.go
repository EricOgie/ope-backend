package dto

type UserDto struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
}
