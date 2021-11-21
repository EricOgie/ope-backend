package service

import (
	"context"
	"net/http"

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
					response.ServeResponse(konstants.ERR, "", res, &ericerrors.EricError{Code: http.StatusUnauthorized, Message: konstants.NO_AUTH})
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
					jwtMapClaim := jwtToken.Claims.(jwt.MapClaims)

					// To enable ushandle CompletePayment flow differently, flow is as follows
					if isCompletePayment(routeInFocus.GetName()) {
						claimObj := models.RetrievePaymentClaim(jwtMapClaim)
						// Embed claim in request context
						ctx := context.WithValue(req.Context(), konstants.DT_KEY, claimObj)
						// send claim through nextHnadler function
						nxtHandler.ServeHTTP(res, req.WithContext(ctx))

					} else {
						claimObj := models.RetrieveClaim(jwtMapClaim)
						// Check Autjorization and respond accordingly
						if authMid.Repo.IsAuthorized(claimObj) == true {
							// Embed claim in request context
							ctx := context.WithValue(req.Context(), konstants.DT_KEY, claimObj)
							// send claim through nextHnadler function
							nxtHandler.ServeHTTP(res, req.WithContext(ctx))

						} else {
							logger.Info("Failed Auth")
							ericErr := ericerrors.NewError(http.StatusForbidden, konstants.UAUTH_ERR)
							response.ServeResponse(konstants.ERR, "", res, ericErr)
						}
					}

				}
			}
		})
	}
}


