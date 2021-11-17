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
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartApp() {

	// define mux router
	router := mux.NewRouter()
	// Load config data
	config := utils.LoadConfig(".")
	// fmt.Println(fmt.Sprintf("%#v", config))
	// Create an instance of DBClient
	dbClient := databases.GetRDBClient(config)
	// Defne a middleware
	midWare := service.AuthMiddlewareService{repositories.MiddleWareRepo{dbClient}}
	// Apply Auth Middleware on router
	router.Use(midWare.AuthMiddleware(config))
	// ------------------------   WIRING AND CONNECTIONS --------------------------
	// userH := handlers.UserHandler{service.NewUserService(repositories.NewUserRepoStub())}
	authH := conhandlers.UserHandler{service.NewUserService(repositories.NewUserRepoDB(dbClient, config))}
	marketH := conhandlers.MarkHandler{service.MarketService{repositories.NewMarketRepoDB(dbClient, config)}}
	fundsH := conhandlers.FundHandler{service.NewPaymentService(repositories.NewFundsRepo(dbClient, config))}

	// Define and include cors handling strategy
	// Cors strategy is currently using a wildcard now. This should change to a selected orrigins when in production
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	credentials := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "PATCH", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling

	// ------------------------   ROUTE DEFINITIONS --------------------------

	// PUBLIC ROUTES
	router.HandleFunc("/", controllers.Greet).Methods(http.MethodGet).Name("Home")
	router.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet).Name("Ping")
	router.HandleFunc("/verify-account", authH.VerifyUserAcc).Methods(http.MethodGet).Name("Verify-Acc")
	router.HandleFunc("/verified", controllers.ServeHTMLTemplate).Methods(http.MethodGet).Name("Verified")
	router.HandleFunc("/request-password-change", authH.RequestPasswordChange).Methods(http.MethodPost).Name("Request-Password-Change")
	router.HandleFunc("/register", authH.CreateUser).Methods(http.MethodPost).Name("RegisterUser")
	router.HandleFunc("/login", authH.Login).Methods(http.MethodPost).Name("Login")

	// - PROTECTED routes
	router.HandleFunc("/users", authH.GetAllUsers).Methods(http.MethodGet).Name("GetAllUser")
	router.HandleFunc("/complete-login", authH.CompleteLoginProcess).Methods(http.MethodPost).Name("Complete-Login")
	router.HandleFunc("/change-password", authH.ChangePassword).Methods(http.MethodPatch).Name("Change-Password")
	router.HandleFunc("/update-profile/{userId:[0-9]+}", authH.UpdateUserProfile).Methods(http.MethodPatch).Name("Profile-Update")
	router.HandleFunc("/user/bankupdate/{userId:[0-9]+}", authH.UpdateUserBank).Methods(http.MethodPatch).Name("Bank-Update")
	router.HandleFunc("/user/refresh", authH.FindOneUser).Methods(http.MethodGet).Name("Refresh")

	router.HandleFunc("/show-market", marketH.FetchMarketState).Methods(http.MethodGet).Name("Market")

	router.HandleFunc("/fund-wallet", fundsH.FundUserWallet).Methods(http.MethodPost).Name("Fund-Wallet")
	router.HandleFunc("/complete-funding", fundsH.CompleteFundingFlow).Methods(http.MethodPatch).Name("Complete-Funding")

	// Start server and log error should ther be one
	logger.Info(konstants.MSG_START + " Address and Port set to " + config.ServerAddress)
	log.Fatal(http.ListenAndServe(":"+config.ServerPort, handlers.CORS(originsOk, headersOk, methodsOk, credentials)(router)))
}
