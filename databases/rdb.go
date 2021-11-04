package databases

import (
	"fmt"
	"time"

	"github.com/EricOgie/ope-be/utils"
	"github.com/jmoiron/sqlx"
)

func GetRDBClient() *sqlx.DB {
	// Get credential setup from environment variables if set
	env, _ := utils.LoadConfig(".")
	// Construct sql connection DATA source
	datasource := fmt.Sprintf("%s@tcp(%s)/%s", env.DBUser, env.DBAddress, env.DBName)
	//Open connection to database
	dbClient, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	// Retrn instance of DB connection
	return dbClient
}
