package serveController

import (
	"fmt"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	response "github.com/EricOgie/ope-be/responses"
)

func Ping(res http.ResponseWriter, req *http.Request) {
	pong := models.Pong{Message: "We Are Live!"}
	response.ServeResponse("Ping-Pong", pong, res, nil)
}

func Greet(res http.ResponseWriter) {
	fmt.Fprintln(res, "<h1>Welcome to Ope service App</h1>")
}
