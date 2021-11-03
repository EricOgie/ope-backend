package response

import (
	"encoding/json"
	"net/http"

	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
)

// ServeResponse serves the correct response to client depending on the result and
//response type preferences. it takes the response resource, statusCode, ResponseWriter,
// and *http.Request

func ServeResponse(collection string, resource interface{},
	res http.ResponseWriter, error *ericerrors.EricError) {

	res.Header().Set(konstants.CONTENT_TYPE, konstants.TYPE_JSON)
	if error != nil {
		res.WriteHeader(error.Code)
		errRes := ErrorResponse{Code: error.Code, Status: "Error", Message: error.Message}
		json.NewEncoder(res).Encode(errRes)

	} else {
		res.WriteHeader(http.StatusOK)
		response := Response{Status: "success", Collection: collection, Data: resource}
		if err := json.NewEncoder(res).Encode(response); err != nil {
			panic(err)
		}
	}

}

type Response struct {
	Status     string      `json:"status" xml:"status"`
	Collection string      `json:"collection" xml:"collection"`
	Data       interface{} `json:"data" xml:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code" xml:"code"`
	Status  string `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}
