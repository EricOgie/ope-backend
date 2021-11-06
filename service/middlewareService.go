package service

import (
	"net/http"
	"strings"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/domain/repositories"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/EricOgie/ope-be/utils"
	"github.com/dgrijalva/jwt-go"
)

// Create an instance of a complete port-port wired AuthMiddleWare.
// i.e, It has a complete to-fro implementation from serviceport layer to RepositoryPort layer
type AuthMiddlewareService struct {
	Repo repositories.MiddleWareRepo
}

/**
* @AUTHHANDLER
* authHnaler function implentaton on AUthMiddleware struct
 */
// AuthMiddleware will validate client authorization/access on all routes that call it usege
func (authMid AuthMiddlewareService) AuthMiddleware(envs utils.Config) func(http.Handler) http.Handler {
	return func(nxtHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

			// Obtain AuthorizationHeader if available
			authorization := req.Header.Get(konstants.AUTH)
			// Hnadle no-auhorization cases
			if authorization == "" {
				logger.Error("EMPTY HEADER")
				response.ServeResponse(
					"Error", "",
					res,
					&ericerrors.EricError{Code: http.StatusUnauthorized, Message: konstants.NO_AUTH})
			} else {
				// Process authoriztion in header
				// pass result to the next handler or abort with a 401 msg
				requestToken := getTokenInHeader(authorization)
				jwtToken, err := jwtTokenFromString(requestToken, envs)

				if err != nil {
					ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
					response.ServeResponse("Error", "", res, ericErr)
				}

				if !jwtToken.Valid {
					logger.Error(konstants.EXP_TOKEN)
					ericErr := ericerrors.NewError(http.StatusForbidden, konstants.EXP_TOKEN)
					response.ServeResponse("Error", "", res, ericErr)
				}

				// Reconstruct Claims from token
				jwtMapClaim := jwtToken.Claims.(jwt.MapClaims)
				claimObj := models.MakeClaim(jwtMapClaim)

				// Check Autjorization and respond accordingly
				if authMid.Repo.IsAuthorized(claimObj) == true {
					nxtHandler.ServeHTTP(res, req)
				} else {
					ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
					response.ServeResponse("Error", "", res, ericErr)
				}
			}

		})
	}
}

// -------------------------- PRIVATE METHODS -------------------------------- //

func getTokenInHeader(header string) string {
	/*
		the header comes in the format below
		Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY...
	*/
	arr := strings.Split(header, " ")
	return arr[1]
}

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
