package controllers

import (
	"net/http"

	serveController "github.com/EricOgie/ope-be/app/serveControllers"
)

// Greet function callable for liveAPP test
func Greet(res http.ResponseWriter, req *http.Request) {
	serveController.Greet(res, req)
}

// Ping should respond with a Pong msg just to hint that application is live
func Ping(res http.ResponseWriter, req *http.Request) {
	serveController.Ping(res, req)
}
