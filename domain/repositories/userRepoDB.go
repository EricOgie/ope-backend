package repositories

import (
	"fmt"
	"time"

	"github.com/EricOgie/ope-be/dto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/setup"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Create UserDB Adapter
type UserRepositoryDB struct {
	client *sqlx.DB
}

// Plug UserRepositoryStub adapter to UserRepository port
func (db UserRepositoryDB) FindAll() (*[]dto.UserDto, *ericerrors.EricError) {

	users := make([]dto.UserDto, 0)
	sqlQuery := "SELECT id, firstname, lastname, email, phone FROM users"

	// Query and marshal to slice of user struct
	err := db.client.Select(&users, sqlQuery)
	// Check error state and responde accordingly
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	return &users, nil
}

// Helper function to instantiate DB
func NewUserRepoDB() UserRepositoryDB {
	// Get credential setup from environment variables if set
	env := setup.GetSetENVs()
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
	return UserRepositoryDB{dbClient}
}
