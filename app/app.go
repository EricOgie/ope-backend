package app

import (
	"log"
	"net/http"

	"github.com/EricOgie/ope-be/app/controllers"
	"github.com/EricOgie/ope-be/app/handlers"
	"github.com/EricOgie/ope-be/domain/repositories"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/service"
	"github.com/EricOgie/ope-be/setup"
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
	env := setup.GetSetENVs()
	logger.Info(konstants.MSG_START + " Address and Port set to " + env.ServerAddress)

	log.Fatal(http.ListenAndServe(env.ServerAddress, router))

}
