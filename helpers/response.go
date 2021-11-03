package helpers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/EricOgie/ope-be/konstants"
)

// ServeResponse serves the correct response to client depending on the response type preferences
// it takes the response resource, statusCode, ResponseWriter, and *http.Request
func ServeResponse(resource interface{}, status int, res http.ResponseWriter, req *http.Request) {

	isXML := req.Header.Get(konstants.CONTENT_TYPE) == konstants.TYPE_XML
	if isXML {
		res.Header().Set(konstants.CONTENT_TYPE, konstants.TYPE_XML)
		if err := xml.NewEncoder(res).Encode(resource); err != nil {
			panic(err)
		}
	} else {
		res.Header().Set(konstants.CONTENT_TYPE, konstants.TYPE_JSON)
		if err := json.NewEncoder(res).Encode(resource); err != nil {
			panic(err)
		}
	}

}
