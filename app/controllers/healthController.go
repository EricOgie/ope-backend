package controllers

import (
	"html/template"
	"net/http"

	serveController "github.com/EricOgie/ope-be/app/serveControllers"
)

// Greet function callable for liveAPP test
func Greet(res http.ResponseWriter, req *http.Request) {
	serveController.Greet(res)
}

// Ping should respond with a Pong msg just to hint that application is live
func Ping(res http.ResponseWriter, req *http.Request) {
	serveController.Ping(res, req)
}

func ServeHTMLTemplate(res http.ResponseWriter, req *http.Request) {
	p := Content{"WELCOME TO GO !", "THIS IS JUST A TEST", "ANOTHER"}
	temp, _ := template.ParseFiles("template.html")
	temp.Execute(res, p)
	// var body bytes.Buffer

	// temp.Execute(&body, p)
}

type Content struct {
	Title    string
	Content1 string
	Content2 string
}
