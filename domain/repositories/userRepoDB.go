package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	_ "github.com/go-sql-driver/mysql"
)

// Create UserDB Adapter
type UserRepositoryDB struct {
	client *sql.DB
}

// Plug UserRepositoryStub adapter to UserRepository port
func (db UserRepositoryDB) FindAll() (*[]models.User, *ericerrors.EricError) {

	sqlQuery := "SELECT id, firstname, lastname, email, phone, created_at FROM users"
	rows, err := db.client.Query(sqlQuery)

	// Should err is not null, return Query error
	if err != nil {
		log.Println("Query Error: " + err.Error())
		return nil, ericerrors.New500Error("Unexpected DB Error")
	}

	// Define user slice and populate with result from query
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password)
		if err != nil {
			ericerrors.New500Error("Unexpected DB Error")
		}

		// Append iteration result to users slice
		users = append(users, user)
	}

	return &users, nil
}

// Helper function to instantiate DB
func NewUserRepoDB() UserRepositoryDB {

	dbClient, err := sql.Open("mysql", "root@tcp(localhost)/ope")
	if err != nil {
		panic(err)
	}

	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return UserRepositoryDB{dbClient}
}
