package requestdto

import (
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// To be called on Loginrequest to validate input
func (req LoginRequest) ValidateRequest() *ericerrors.EricError {

	if !isValidEmail(req.Email) {
		return ericerrors.New422Error(konstants.INVALID_EMAIL)
	}

	if !isValidPword(req.Password) {
		return ericerrors.New422Error(konstants.INVALID_PWORD)
	}

	return nil
}
