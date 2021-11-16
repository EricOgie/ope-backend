package service

import (
	"net/http"
	"strings"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/logger"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/utils"
	"github.com/dgrijalva/jwt-go"
)

// This will check the active or in focus route if listed among the authenticatable routes
func needsAuthorization(routeName string) bool {
	auth := map[string]bool{
		"Home":                    false,
		"Ping":                    false,
		"Verify-Acc":              false,
		"Verified":                false,
		"Login":                   false,
		"Request-Password-Change": false,
		"RegisterUser":            false,
		"GetAllUser":              true,
		"Complete-Login":          true,
		"Change-Password":         true,
		"Market":                  true,
		"Fund-Wallet":             true,
		"Complete-Funding":        true,
	}
	return auth[routeName]

}

func isCompletePayment(urlName string) bool {
	return urlName == "Complete-Funding"
}

func isTokenInURL(req *http.Request) bool {
	tok := req.URL.Query().Get("k")
	return len(tok) > 0
}

func getTokenInHeader(header string) string {
	/*
		the header comes in the format below
		Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY...
	*/
	arr := strings.Split(header, " ")
	return arr[1]
}

// This converts the passed jwtstring into twt.token
func jwtTokenFromString(tokenString string, ens utils.Config) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(ens.SigningKey), nil
	})

	if err != nil {
		logger.Error("PARSE ERROR e= " + err.Error())
		return nil, err
	}
	return token, nil
}

func getClaim(req *http.Request, env utils.Config, res http.ResponseWriter) models.Claim {
	tokString := req.URL.Query().Get("k")
	// convert to JWT
	tokJwt, e := jwtTokenFromString(tokString, env)
	if e != nil {
		logger.Error("JWT EXTRACT ERR : " + e.Error())
	}

	if !tokJwt.Valid {
		response.ServeResponse("Error", "", res,
			&ericerrors.EricError{Code: http.StatusUnauthorized, Message: "Expired Authorization"})
	}

	jwtMapClaim := tokJwt.Claims.(jwt.MapClaims)
	claimObj := models.RetrieveClaim(jwtMapClaim)
	return claimObj

}
