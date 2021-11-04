package app

import (
	"log"
	"net/http"

	"github.com/EricOgie/ope-be/app/controllers"
	"github.com/EricOgie/ope-be/app/handlers"
	"github.com/EricOgie/ope-be/databases"
	"github.com/EricOgie/ope-be/domain/repositories"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/service"
	"github.com/EricOgie/ope-be/utils"
	"github.com/gorilla/mux"
)

func StartApp() {

	// define mux router
	router := mux.NewRouter()
	// Load config data
	config, err := utils.LoadConfig(".")
	// Create an instance of DBClient
	dbClient := databases.GetRDBClient()
	// Sanity Check
	utils.RunSanityCheck(err)

	// ------------------------   WIRING AND CONNECTIONS --------------------------
	// userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoStub())}
	userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoDB(dbClient))}

	// ------------------------   ROUTE DEFINITIONS --------------------------

	// Health check routs
	router.HandleFunc("/", controllers.Greet).Methods(http.MethodGet)
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)

	// User related routes
	router.HandleFunc("/users", userH.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/register-users", userH.CreateUser).Methods(http.MethodPost)

	// Start server and log error should ther be one
	logger.Info(konstants.MSG_START + " Address and Port set to " + config.ServerAddress)

	log.Fatal(http.ListenAndServe(config.ServerAddress, router))

}
