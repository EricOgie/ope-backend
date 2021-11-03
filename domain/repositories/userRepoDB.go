package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	_ "github.com/go-sql-driver/mysql"
)

// Create UserDB Adapter
type UserRepositoryDB struct {
	client *sql.DB
}

// Plug UserRepositoryStub adapter to UserRepository port
func (db UserRepositoryDB) FindAll() ([]models.User, error) {

	sqlQuery := "SELECT id, firstname, lastname, email, phone, created_at FROM users"
	rows, err := db.client.Query(sqlQuery)

	if err != nil {
		log.Println("Query Error: " + err.Error())
		return nil, err
	}
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password)
		if err != nil {
			log.Println("Scan Error: " + err.Error())
		}

		users = append(users, user)
	}

	return users, err
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
