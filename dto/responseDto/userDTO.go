package responsedto

// One UserDto for a single user response
type OneUserDto struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
	CreatedAt string `db:"created_at"`
	Token     string `json:"token" xml:"token"`
}

// User DTO for a multiple user response
type UserDto struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `json:"updated_at" xml:"updated_at"`
}
