package repositories

import (
	"strconv"

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
	id, _ := strconv.Atoi(claim.Id)

	querySQL := "SELECT id, firstname, lastname, email, phone, created_at FROM users WHERE id = ?"
	var user models.User
	err := repo.Client.Get(&user, querySQL, id)
	// Check error state and respond accordingly
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
	if c.CreatedAt == u.CreatedAt && c.Id == u.Id {
		return true

	} else {
		return false
	}
}
