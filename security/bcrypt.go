package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPword(pword string) (string, error) {
	hashedPword, err := bcrypt.GenerateFromPassword([]byte(pword), 14)
	return string(hashedPword), err
}

func CheckUserPassword(password, hashedPword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPword), []byte(password))
	return err == nil
}
