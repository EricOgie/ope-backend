package responsedto

// One UserDto for a single user response
type OneUserDto struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
	CreatedAt string `db:"created_at" json:"created_at"`
	Token     string `json:"token" xml:"token"`
}

type CompleteUserDTO struct {
	Id          string `json:"user_id"`
	FirstName   string `json:"firstname" xml:"first_name"`
	LastName    string `json:"lastname" xml:"last_name"`
	Email       string `json:"email" xml:"email"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	Holdings    string
	BankAccount BankAccountDTO `json:"bank_account"`
	Otp         string         `json:"otp"`
	Wallet      WalletDTO      `json:"wallet"`
	Token       string         `json:"token" xml:"token"`
	Portfolio   interface{}    `json:"portfolio" xml:"portfolio"`
}

type UserProfileDTO struct {
	Id          string         `json:"user_id"`
	FirstName   string         `json:"firstname" xml:"first_name"`
	LastName    string         `json:"lastname" xml:"last_name"`
	Email       string         `json:"email" xml:"email"`
	Phone       string         `json:"phone" xml:"phone"`
	BankAccount BankAccountDTO `json:"bank_account"`
}

type OneUserDtoWithOtp struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
	OTP       int    `json:"otp" xml:"otp"`
	CreatedAt string `db:"created_at" json:"created_at"`
	Token     string `json:"token" xml:"token"`
}

// User DTO for a multiple user response
type UserDto struct {
	Id        string `json:"user_id"`
	FirstName string `json:"firstname" xml:"first_name"`
	LastName  string `json:"lastname" xml:"last_name"`
	Email     string `json:"email" xml:"email"`
	Phone     string `json:"phone" xml:"phone"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `json:"updated_at" xml:"updated_at"`
}

type VerifiedRESPONSE struct {
	Id     string `json:"user_id"`
	Email  string `json:"email" xml:"email"`
	Status string `json:"status" xml:"status"`
}

type LoginResponseDTO struct {
	Message     string `json:"message"`
	TokenString string `json:"token"`
}

type PlainResponseDTO struct {
	Code    int
	Message string
}

func (u CompleteUserDTO) GetOneUserFromComplete() OneUserDto {
	return OneUserDto{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		Token:     u.Token,
	}
}
