package repositories

import (
	"database/sql"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Create UserDB Adapter
type UserRepositoryDB struct {
	client *sqlx.DB
}

// Helper function to instantiate DB
func NewUserRepoDB(dbClient *sqlx.DB) UserRepositoryDB {
	return UserRepositoryDB{dbClient}
}

/**
* @FINDALL
* METHOD implemetation of UserRepositoryPort as an interface
* Interface implementation here correspond to Plugging  UserRepositoryDB
* adapter to UserRepositoryPort
* Only callable when active user is an ADMIN
 */

func (db UserRepositoryDB) FindAll() (*[]responsedto.UserDto, *ericerrors.EricError) {

	users := make([]responsedto.UserDto, 0)
	sqlQuery := "SELECT id, firstname, lastname, email, phone, created_at FROM users"
	// Query and marshal to slice of user struct
	err := db.client.Select(&users, sqlQuery)
	// Check error state and responde accordingly
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	return &users, nil
}

/**
* @CREATE
* METHOD implemetation of UserRepositoryPort as an interface
* To be called upon REGISTER user Request
 */
func (db UserRepositoryDB) Create(u models.User) (*models.User, *ericerrors.EricError) {
	// Define Query
	insertQuery := "INSERT INTO users (firstname, lastname, email, phone, password, created_at) " +
		"values(?, ?, ?, ?, ?, ?)"
	// Hash User password
	hashedPword := security.GenHashedPwd(u.Password)
	// Execute query
	result, err := db.client.Exec(insertQuery, u.FirstName,
		u.LastName, u.Email, u.Phone, hashedPword, u.CreatedAt)

	// Handle possible Error
	if err != nil {
		logger.Error(konstants.DB_INSERT_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// Ubtain last insertedID
	newId, err := result.LastInsertId()
	// Handle possible ID retrieval  Error
	if err != nil {
		logger.Error(konstants.DB_ID_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// merge ID from Db with UserObject
	u.Id = strconv.Itoa(int(newId))
	// return newly created user
	return &u, nil

}

/**
* @LOGIN
* To be called upon Login Request
 */

func (db UserRepositoryDB) Login(u models.UserLogin) (*models.User, *ericerrors.EricError) {
	if !isValidCrentials(u, db) {
		// User either does not exist or Email-Password does not match
		logger.Error(konstants.CREDENTIAL_ERR)
		return nil, ericerrors.NewCredentialError(konstants.CREDENTIAL_ERR)
	}
	// Define Query
	querySQL := "SELECT id, firstname, lastname, email, phone, created_at FROM users WHERE email = ?"
	var user models.User
	err := db.client.Get(&user, querySQL, u.Email)
	// Check error state and responde accordingly
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}
	return &user, nil
}

// ---------------------- PRIVATE METHODS ------------------------//
func isValidCrentials(u models.UserLogin, client UserRepositoryDB) bool {
	stm := "SELECT password FROM users WHERE email = ?"
	data := ""
	err := client.client.Get(&data, stm, u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(konstants.NO_DER_ERR + err.Error())
		} else {
			logger.Error(konstants.DB_ERROR + err.Error())
		}
	}
	return security.CheckUserPassword(u.Password, data)
}
