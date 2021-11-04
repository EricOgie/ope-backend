package repositories

import (
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
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

// Interface implementation here correspond to Plugging  UserRepositoryDB adapter to UserRepository p
// Find all Users only callable when active user is an ADMIN
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

// Register a New User
func (db UserRepositoryDB) Create(u models.User) (*models.User, *ericerrors.EricError) {
	// Define Query
	insertQuery := "INSERT INTO users (firstname, lastname, email, phone, password, created_at) " +
		"values(?, ?, ?, ?, ?, ?)"
		// Execute query
	result, err := db.client.Exec(insertQuery, u.FirstName,
		u.LastName, u.Email, u.Phone, u.Password, u.CreatedAt)

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
