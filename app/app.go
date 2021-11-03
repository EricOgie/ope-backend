package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EricOgie/ope-be/app/controllers"
	"github.com/EricOgie/ope-be/app/handlers"
	"github.com/EricOgie/ope-be/domain/repositories"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/service"
	"github.com/gorilla/mux"
)

func StartApp() {
	// define mux router
	router := mux.NewRouter()

	// ------------------------   WIRING AND CONNECTIONS --------------------------
	// userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoStub())}
	userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoDB())}

	// ------------------------   ROUTE DEFINITIONS --------------------------

	// Health check routs
	router.HandleFunc("/", controllers.Greet).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)

	// User related routes
	router.HandleFunc("/users", userH.GetAllUsers).Methods(http.MethodGet)

	// Start server and log error should ther be one
	fmt.Println(konstants.MSG_START)
	log.Fatal(http.ListenAndServe(konstants.LOCAL_ADD, router))

}
