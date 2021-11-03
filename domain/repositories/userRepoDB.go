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

	if err != nil {
		log.Println("Query Error: " + err.Error())
		return nil, ericerrors.New404Error("No Resource Found")
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password)

		if err != nil {
			// Two error cases can result in this case; 500 and 404 type errors
			// We will check for a 404 type
			if err == sql.ErrNoRows {
				return nil, ericerrors.New404Error("No Resource Found")
			} else {
				log.Println("Scan Error: " + err.Error())
				return nil, ericerrors.New500Error("Unexpected Server Error")
			}

		}

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
