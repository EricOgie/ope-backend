package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EricOgie/ope-be/app/conhandlers"
	"github.com/EricOgie/ope-be/app/controllers"
	"github.com/EricOgie/ope-be/databases"
	"github.com/EricOgie/ope-be/domain/repositories"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/service"
	"github.com/EricOgie/ope-be/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartApp() {

	// define mux router
	router := mux.NewRouter()
	// Load config data
	config := utils.LoadConfig(".")
	fmt.Println(fmt.Sprintf("%#v", config))
	// Create an instance of DBClient
	dbClient := databases.GetRDBClient(config)
	// Defne a middleware
	midWare := service.AuthMiddlewareService{repositories.MiddleWareRepo{dbClient}}
	// Apply Auth Middleware on router
	router.Use(midWare.AuthMiddleware(config))
	// ------------------------   WIRING AND CONNECTIONS --------------------------
	// userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoStub())}
	authH := conhandlers.UserHandler{service.NewUserService(repositories.NewUserRepoDB(dbClient, config))}

	// ------------------------   ROUTE DEFINITIONS --------------------------
	// port := os.Getenv("PORT")
	// PUBLIC ROUTES
	router.HandleFunc("/", controllers.Greet).Methods(http.MethodGet).Name("Home")
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet).Name("Ping")
	router.HandleFunc("/verify", controllers.ServeHTMLTemplate).Methods(http.MethodGet).Name("Verify")
	router.HandleFunc("/register", authH.CreateUser).Methods(http.MethodPost).Name("RegisterUser")
	router.HandleFunc("/login", authH.Login).Methods(http.MethodPost).Name("Login")

	// - PROTECTED routes
	router.HandleFunc("/users", authH.GetAllUsers).Methods(http.MethodGet).Name("GetAllUser")
	router.HandleFunc("/complete-login", authH.CompleteLoginProcess).Methods(http.MethodPost).Name("Complete-Login")

	// Start server and log error should ther be one
	// Define and include cors handling strategy
	// Cors strategy is currently using a wild now This should change to a of selected orrigins when in production
	logger.Info(konstants.MSG_START + " Address and Port set to " + config.ServerAddress)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort,
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))

}
