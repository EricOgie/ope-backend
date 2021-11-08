package repositories

import (
	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/jmoiron/sqlx"
)

type MiddleWarePort interface {
	IsAuthorized(claim models.Claim) bool
}

type MiddleWareRepo struct {
	Client *sqlx.DB
}

func (repo MiddleWareRepo) IsAuthorized(claim models.Claim) bool {
	querySQL := "SELECT id, firstname, lastname, email, phone, created_at FROM users WHERE email = ?"
	var user models.User
	err := repo.Client.Get(&user, querySQL, claim.Email)
	// Check error state and responde accordingly
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return false
	} else {
		// Check If name, createdAt and phone mathes same from DB
		return isValidAuth(claim, user)
	}
}

//  ----------------------------- PRIVATE METHOD ------------------------ //

func isValidAuth(c models.Claim, u models.User) bool {
	if c.Firstname == u.FirstName && c.Lastname == u.LastName && c.When == u.CreatedAt {

		return true
	} else {

		return false
	}
}
