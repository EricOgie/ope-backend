package databases

import (
	"fmt"
	"time"

	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/jmoiron/sqlx"
)

// GetRDBClient establishes a RDB connection and return a single instane of an *sqlx.DB.
// It takes utils.Config struct as input.
// The utils.Config struct intake here, prevent having to reload and READ env variables
func GetRDBClient(env utils.Config) *sqlx.DB {
	// Construct sql connection DATA source
	r := env.DBUser + ":" + env.DBPassword + "@tcp(" + env.DBAddress + ":" + env.DBPort + ")" + "/" + env.DBName
	logger.Info("datasource = " + r)
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DBUser, env.DBPassword, env.DBAddress, env.DBPort, env.DBName)
	//Open connection to database
	dbClient, err := sqlx.Open("mysql", datasource)
	if err != nil {
		logger.Error(konstants.DB_CON_ERR + err.Error())
		panic(err)
	} else {
		// LOG  Con Success
		logger.Info(konstants.DB_CON_OK)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	// Retrn instance of DB connection
	return dbClient
}
