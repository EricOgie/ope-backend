package app

import (
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

	// "github.com/gorilla/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartApp() {

	// Define cors handling strategy
	// head := handlers.AllowedHeaders([]string{"X-Requested-With", "Content_Type", "Authorization"})
	// mtd := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "PUT"})
	// origin := handlers.AllowedOrigins(([]string{"*"}))

	options := []string{"*", "http://localhost:*", "http://localhost:8080/", "localhost:8080/", "https://loaner-two.vercel.app/"}
	cors := handlers.CORS(handlers.AllowedOrigins(options))

	// define mux router
	router := mux.NewRouter()
	// Load config data
	config := utils.LoadConfig(".")

	// Create an instance of DBClient
	dbClient := databases.GetRDBClient(config)
	// Create an instance of SMTPClient that will be use for mailing
	// This way, we don get to create multiple smtp connections because we just
	// have one instance and pass it along when and where it is needed
	// smptClient := utils.GetEmailClient(config)

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
	logger.Info(konstants.MSG_START + " Address and Port set to " + config.ServerAddress)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, cors(router)))

}
