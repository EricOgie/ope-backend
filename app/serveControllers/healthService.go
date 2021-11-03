package serveController

import (
	"fmt"
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/helpers"
)

func Ping(res http.ResponseWriter, req *http.Request) {
	pong := models.Pong{Message: "We Are Live!"}
	helpers.ServeResponse(pong, res, nil)
}

func Greet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "<h1>Welcome to Ope service App</h1>")
}
