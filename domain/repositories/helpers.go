package repositories

import (
	"net/http"
	"strconv"
	"time"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
)

// ---------------------- PRIVATE METHODS ------------------------//

func runUserQueryWithEmail(userEmail string, db UserRepositoryDB) (*models.CompleteUser, *ericerrors.EricError) {
	// querySQL := "SELECT id, firstname, lastname, email, phone, password, created_at FROM users WHERE email = ?"
	querySQL := "SELECT users.id, users.firstname, users.lastname, users.email, users.phone, users.password," +
		" users.created_at, users.account_no, users.account_name, wallet.amount, wallet.address FROM wallet" +
		" INNER JOIN users ON wallet.user_id = users.id WHERE wallet.user_id = ?"
	var user models.QueryUser
	userId := UserId(userEmail, db)

	err := db.client.Get(&user, querySQL, userId)
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
	allInOne := user.MakeCompleteUserFromQueryUser()
	return &allInOne, nil
}

//
func userIsRegistered(userEmail string, db UserRepositoryDB) bool {
	querySQL := "SELECT  email FROM users WHERE email = ?"
	var user models.User
	err := db.client.Get(&user, querySQL, userEmail)
	return err == nil
}

func createUserWallet(db UserRepositoryDB, firstname string, userId int64) (*models.Wallet, *ericerrors.EricError) {
	sqlQ := "INSERT INTO wallet (amount, address, user_id) VALUES (?, ?, ?)"
	walletAdd := genWalletAddress(firstname)

	res, err := db.client.Exec(sqlQ, 0.00, walletAdd, userId)
	if err != nil {
		logger.Error("Wallet Err: " + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// Ubtain last insertedID
	newId, err := res.LastInsertId()
	// Handle possible ID retrieval  Error
	if err != nil {
		logger.Error(konstants.DB_ID_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	newWallet := models.Wallet{Id: newId, Amount: 0.00, Address: walletAdd}
	return &newWallet, nil

}

func genWalletAddress(firstName string) string {
	key := firstName + time.Now().String()
	walletadd, err := security.Hash(key)
	if err != nil {
		logger.Error("Hash Err: " + err.Error())
		return ""
	}

	return walletadd

}

func UserId(email string, db UserRepositoryDB) int32 {
	querySQL := "SELECT  id FROM users WHERE email = ?"
	var userid int32
	err := db.client.Get(&userid, querySQL, email)
	if err != nil {
		logger.Error("EROOROOO" + err.Error())
	}

	logger.Info("USERID = " + strconv.Itoa(int(userid)))
	return userid
}
