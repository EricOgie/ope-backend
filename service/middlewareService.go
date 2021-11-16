package service

import (
	"context"
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
	"github.com/gorilla/mux"
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

			routeInFocus := mux.CurrentRoute(req)

			if !needsAuthorization(routeInFocus.GetName()) {
				// Check if token in url, It might just be a case of email verification
				if isTokenInURL(req) {
					// extract claim data and pass along with request
					claim := getClaim(req, envs, res)
					ctx := context.WithValue(req.Context(), konstants.DT_KEY, claim)
					nxtHandler.ServeHTTP(res, req.WithContext(ctx))
				} else {
					nxtHandler.ServeHTTP(res, req)
				}

			} else {

				// Obtain AuthorizationHeader if available
				authorization := req.Header.Get(konstants.AUTH)
				// Hnadle no-auhorization cases
				if authorization == "" {
					logger.Error("EMPTY HEADER")
					response.ServeResponse(
						konstants.ERR, "", res,
						&ericerrors.EricError{Code: http.StatusUnauthorized, Message: konstants.NO_AUTH})
				} else {
					// Process authoriztion in header
					// pass result to the next handler or abort with a 401 msg
					requestToken := getTokenInHeader(authorization)
					jwtToken, err := jwtTokenFromString(requestToken, envs)

					if err != nil {
						ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
						response.ServeResponse(konstants.ERR, "", res, ericErr)
						return
					}

					if !jwtToken.Valid {
						logger.Error(konstants.EXP_TOKEN)
						ericErr := ericerrors.NewError(http.StatusForbidden, konstants.EXP_TOKEN)
						response.ServeResponse(konstants.ERR, "", res, ericErr)
					}

					// Reconstruct Claims from token
					jwtMapClaim := jwtToken.Claims.(jwt.MapClaims)
					claimObj := models.MakeClaim(jwtMapClaim)

					// Check Autjorization and respond accordingly
					if authMid.Repo.IsAuthorized(claimObj) == true {
						// Embed claim in request context
						ctx := context.WithValue(req.Context(), konstants.DT_KEY, claimObj)
						// send claim through nextHnadler function
						nxtHandler.ServeHTTP(res, req.WithContext(ctx))

					} else {
						logger.Info("ERERE")
						ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
						response.ServeResponse(konstants.ERR, "", res, ericErr)
					}
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
	}
	return auth[routeName]

}

func isTokenInURL(req *http.Request) bool {
	tok := req.URL.Query().Get("k")
	return len(tok) > 0
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
	claimObj := models.MakeClaim(jwtMapClaim)
	return claimObj

}
