package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
)

type AuthorizationPort interface {
	IsAuthorized(token string, routeName string, reqVars map[string]string) bool
}

type AuthorizationRepo struct{}

func (authRepo AuthorizationRepo) IsAuthorized(token string, routeName string, reqVars map[string]string) bool {
	// Generate  encoded rul string using token, routeName and reqVars
	url := generateURL(token, routeName, reqVars)
	// call to server
	res, err := http.Get(url)

	// If error is not nil, caller should get a false boolean resp
	// otherwise, the flow runs on to construct the resp into a map
	if err != nil {
		logger.Error(konstants.ERR_REQ_SEND + err.Error())
		return false
	} else {
		// Construct response into json object
		m := map[string]bool{}
		err := json.NewDecoder(res.Body).Decode(&m)
		if err != nil {
			logger.Error(konstants.ERR_DECODE + err.Error())
			return false
		}
		// return the value of the key, isAuthorized
		return m["isAuthorized"]
	}
}

// -------------------------------- PRIVATE METHODS ---------------------------------- //

// Call to generate an encoded string url using the specified args/params
func generateURL(token string, routeName string, vars map[string]string) string {
	config, _ := utils.LoadConfig(".")
	u := url.URL{Host: config.ServerAddress, Path: "/auth", Scheme: "http"}
	query := u.Query()
	query.Add("token", token)
	query.Add("routeName", routeName)

	for key, val := range vars {
		query.Add(key, val)
	}

	u.RawQuery = query.Encode()
	return u.String()
}
