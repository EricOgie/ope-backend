package security

import (
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
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

func GenHashedPwd(pword string) string {
	pWord, err := HashPword(pword)
	if err != nil {
		logger.Error(konstants.HASH_ERR + err.Error())
		return ""
	} else {
		return pWord
	}
}
