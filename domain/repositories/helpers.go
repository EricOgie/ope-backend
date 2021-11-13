package repositories

import (
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
)

// ---------------------- PRIVATE METHODS ------------------------//

func runUserQueryWithEmail(userEmail string, db UserRepositoryDB) (*models.User, *ericerrors.EricError) {
	querySQL := "SELECT id, firstname, lastname, email, phone, password, created_at FROM users WHERE email = ?"
	var user models.User

	err := db.client.Get(&user, querySQL, userEmail)
	// Check error state and responde accordingly
	if err != nil {
		if err.Error() == konstants.DB_NO_ROW {
			// user does not exist
			logger.Error(konstants.DB_ERROR + konstants.CREDENTIAL_ERR)
			return nil, ericerrors.NewError(http.StatusUnauthorized, konstants.CREDENTIAL_ERR)
		} else {
			logger.Error(konstants.QUERY_ERR + err.Error())
			return nil, ericerrors.New500Error(konstants.MSG_500)
		}
	}
	return &user, nil
}

func userIsRegistered(userEmail string, db UserRepositoryDB) bool {
	querySQL := "SELECT  email FROM users WHERE email = ?"
	var user models.User
	err := db.client.Get(&user, querySQL, userEmail)
	return err == nil
}
