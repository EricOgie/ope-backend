package repositories

import (
	"net/http"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/security"
	"github.com/EricOgie/ope-be/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Create UserDB Adapter
type UserRepositoryDB struct {
	client *sqlx.DB
}

// Helper function to instantiate DB
func NewUserRepoDB(dbClient *sqlx.DB, env utils.Config) UserRepositoryDB {
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
func (db UserRepositoryDB) Create(u models.User) (*models.CompleteUser, *ericerrors.EricError) {
	// First check if User is registered prior
	if userIsRegistered(u.Email, db) {
		logger.Info("User with email, " + u.Email + " is registered prior ")
		return nil, &ericerrors.EricError{Code: 403, Message: konstants.MSG_403}
	}
	// Define Query
	insertQuery := "INSERT INTO users (firstname, lastname, email, phone, password, created_at) " +
		"values(?, ?, ?, ?, ?, ?)"
	// Hash User password
	logger.Info("PASS=" + u.Password)
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

	// Create wallet for user
	wallet, ericErr := createUserWallet(db, u.FirstName, newId)
	if ericErr != nil {
		logger.Error("Wallet Err: " + ericErr.Message)
	}

	// merge ID from Db with UserObject
	u.Id = strconv.Itoa(int(newId))
	// construct user with wallet and bank struck
	userWithWallet := u.MakeCompleteUser(wallet)
	// return newly created user
	return &userWithWallet, nil

}

/**
* @LOGIN
* To be called upon Login Request
 */

func (db UserRepositoryDB) Login(u models.UserLogin) (*models.CompleteUser, *ericerrors.EricError) {
	return runUserQueryWithEmail(u.Email, db)
}

/**
* @VERIFYUSERACCOUNT
* METHOD implemetation of UserRepositoryPort as an interface
* To be called upon verification of user
 */

func (db UserRepositoryDB) VerifyUserAccount(v models.VerifyUser) (*models.User, *ericerrors.EricError) {
	query := "UPDATE users SET is_verified = ? WHERE email = ?"
	_, err := db.client.Exec(query, "true", v.Email)
	if err != nil {
		logger.Error(konstants.VET_ACC_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	user := v.GetUserFromVerify()
	return &user, nil

}

/**
* @CHANGEPASSWORD
* METHOD implemetation of UserRepositoryPort as an interface
* To be called to commplete login workflow
 */
func (db UserRepositoryDB) CompleteLogin(claim models.Claim) (*models.CompleteUser, *ericerrors.EricError) {
	userId, err := strconv.Atoi(claim.Id)

	if err != nil {
		logger.Error(konstants.STRING_INT_ERR + err.Error())
	}
	sqlQuery := "SELECT id, symbol, image, total_quantity, unit_price, equity_value, percentage_change FROM stocks WHERE user_id = ?"
	userStocks := make([]models.Stock, 0)
	// Query and marshal to slice of stock-struct
	qErr := db.client.Select(&userStocks, sqlQuery, userId)

	// Handle possible query error
	if qErr != nil {
		logger.Error(konstants.QUERY_ERR + qErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// Cretate a completeuser struct by merging the user struct in the claim passed into function with
	// the slice of user stocks gotten from the DB. This will serve a user with his/her stock portfolio
	completeUser := models.MakeCompleteUser(claim, userStocks)
	return &completeUser, nil
}

/**
* @REQUESTPASSWORDCHANGE
* METHOD implemetation of UserRepositoryPort as an interface
* To be called upon password change request
 */
func (db UserRepositoryDB) RequestPasswordChange(userEmail models.UserEmail) (*models.CompleteUser, *ericerrors.EricError) {
	return runUserQueryWithEmail(userEmail.Email, db)
}

/**
* @CHANGEPASSWORD
* METHOD implemetation of UserRepositoryPort as an interface
* To be called to commplete password change workflow
 */
func (db UserRepositoryDB) ChangePassword(u models.UserLogin) (*responsedto.PlainResponseDTO, *ericerrors.EricError) {
	hashedPword := security.GenHashedPwd(u.Password)
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := db.client.Exec(query, hashedPword, u.Email)
	if err != nil {
		logger.Error("Edit Error" + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	res := u.GetPlainResponseDTO(http.StatusOK, "Password Changed")
	return &res, nil
}
