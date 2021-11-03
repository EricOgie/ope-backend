package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
)

// ServeResponse serves the correct response to client depending on the response type preferences
// it takes the response resource, statusCode, ResponseWriter, and *http.Request
func ServeResponse(resource interface{}, res http.ResponseWriter, error *ericerrors.EricError) {
	data := map[string]interface{}{"status": "success", "type": "User", "data": resource}
	res.Header().Set(konstants.CONTENT_TYPE, konstants.TYPE_JSON)
	if error != nil {
		res.WriteHeader(error.Code)
		fmt.Fprintf(res, error.Message)
	}

	if err := json.NewEncoder(res).Encode(data); err != nil {
		panic(err)
	}

}
