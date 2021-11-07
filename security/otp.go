package security

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// GenerateOPT is a callable func that output 6-digit int
// Should be called at the point of Uster struct conversion to OneUserDTO
func GenerateOTP() int {
	otpString := ""
	for i := 0; i < 6; i++ {
		opeRand, _ := rand.Int(rand.Reader, big.NewInt(9))
		otpString += opeRand.String()
	}

	otp, _ := strconv.Atoi(otpString)
	return otp

}
