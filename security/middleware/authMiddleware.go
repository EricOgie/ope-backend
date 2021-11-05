package middleware

import (
	"net/http"
	"strings"

	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	response "github.com/EricOgie/ope-be/responses"
	"github.com/gorilla/mux"
)

// Define a middleware struct and inject user or auth repository
type AuthMiddleware struct {
	repo AuthorizationPort
}

/**
* @AUTHHANDLER
* authHnaler function implentaton on AUthMiddleware struct
 */
func (auth AuthMiddleware) authHandler() func(http.Handler) http.Handler {
	return func(nxtHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			routeInfocus := mux.CurrentRoute(req)
			// Get route variable from request
			routeVars := mux.Vars(req)
			// Obtain AuthorizationHeader if available
			authorization := req.Header.Get(konstants.AUTH)

			// Hnadle no-auhorization cases
			if authorization == "" {
				response.ServeResponse(
					"Error", "",
					res,
					&ericerrors.EricError{Code: http.StatusUnauthorized, Message: konstants.NO_AUTH})
			}

			// Process authoriztion in header
			// pass result to the next handler or abort with a 403 msg
			requestToken := getTokenInHeader(authorization)
			isAuthorizedValue := auth.repo.IsAuthorized(requestToken, routeInfocus.GetName(), routeVars)

			if isAuthorizedValue {
				nxtHandler.ServeHTTP(res, req)
			} else {
				ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
				response.ServeResponse("Error", "", res, ericErr)
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
