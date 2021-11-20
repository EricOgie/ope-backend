package repositories

import (
	"net/http"

	"github.com/EricOgie/ope-be/domain/models"
	responsedto "github.com/EricOgie/ope-be/dto/responseDto"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
	"github.com/EricOgie/ope-be/utils"
	"github.com/jmoiron/sqlx"
)

type LoanRepo struct {
	Client *sqlx.DB
}

// DB injector
func NewLoanRepo(client *sqlx.DB, env utils.Config) LoanRepo {
	return LoanRepo{Client: client}
}

/**
* @TAKELOAN
* METHOD implemetation of LoanRepositoryPort as an interface
* Interface implementation here correspond to Plugging  UserRepositoryDB
* adapter to UserRepositoryPort
* Callable n service layer to register a loan
 */

func (db LoanRepo) TakeLoan(loan models.Loan) (*responsedto.LoanResDTO, *ericerrors.EricError) {
	// First check that loan amount is <= 60% of user Portfolior position
	isPassed60PercentCheck := Check60PercentMark(db, loan.Amount, loan.UserId)

	if !isPassed60PercentCheck {
		logger.Error("Failed 60% Check")
		eErro := ericerrors.NewError(http.StatusNotAcceptable, "Loan can not exceed 60% of total Investment")
		return nil, eErro
	}
	//Prapare SQL statement
	query := "INSERT INTO loans (user_id, amount, package, duration) values(?, ?, ?, ?)"

	result, qErr := db.Client.Exec(query, loan.UserId, loan.Amount, loan.Package, loan.Duration)
	// Hndle Error ase
	if qErr != nil {
		logger.Error(konstants.QUERY_ERR + qErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	loanId, err := result.LastInsertId()

	if err != nil {
		logger.Error(konstants.DB_ID_ERR + err.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// Use last inserted id as loan id
	loan.Id = int(loanId)
	response := loan.ConvertToLoanResDTO()
	return &response, nil

}

//
func (db LoanRepo) FetchLoans(userId int) (*[]models.QueryLoan, *ericerrors.EricError) {
	query := "SELECT id, amount, paid, package, duration, status, created_at FROM loans WHERE user_id = ?"

	loans := make([]models.QueryLoan, 0)
	qErr := db.Client.Select(&loans, query, userId)

	if qErr != nil {
		logger.Error(konstants.QUERY_ERR + qErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	return &loans, nil
}

func (db LoanRepo) PayLoan(pay models.LoanPayment) (*responsedto.RepaymentResDTO, *ericerrors.EricError) {
	// Check user wallet sufficency for loan installment
	isSufficient := CheckWallet(db, pay.Payment, pay.UserId)
	if !isSufficient {
		logger.Error(konstants.ERR_INSURFICIENCY)
		return nil, ericerrors.NewError(http.StatusNotAcceptable, konstants.ERR_INSURFICIENCY)
	}
	// Get current state of loan
	loan, loanErr := getLoan(pay.LoanId, db)
	if loanErr != nil {
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}
	// Check if loan is open or close and respond accordingly
	if loan.Status != konstants.LOAN_OPEN {
		ericErr := ericerrors.NewError(http.StatusNotAcceptable, konstants.LOAN_INACTIVE)
		return nil, ericErr
	}

	// calculate correct payment from user past transactions so user don't send in more than is required
	correctPay := getCorrectPaymnet(loan, pay)
	balanceAfterPay := loan.Amount - (loan.Paid + correctPay)
	// Less Amount from user's wallet and add pay loan installment
	walError := minusFromWallet(db, pay.UserId, correctPay)
	if walError != nil {
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}
	// record installment
	query := "INSERT INTO repayments (payment, balance, created_at, loan_id, user_id) values (?, ?, ?, ?, ?)"

	result, qErr := db.Client.Exec(query, correctPay, balanceAfterPay, pay.CreatedAt, loan.Id, pay.UserId)
	if qErr != nil {
		logger.Error("DEDUCTED BUT NO REC. Err: " + qErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}
	// Update Loan state
	loanStatsu, uErr := updateLoan(db, loan, correctPay)
	if uErr != nil {
		logger.Error("Loan Update Err" + uErr.Message)
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	// make RepaymentResDTO and return
	payId, _ := result.LastInsertId()
	repayDTO := pay.ConvertToRepaymentDTO(correctPay, int(payId), loan, loanStatsu)

	return &repayDTO, nil

}

func (db LoanRepo) GetInstallments(loanId int) (*[]models.Querypayment, *ericerrors.EricError) {

	installments := make([]models.Querypayment, 0)

	query := "SELECT repayments.id, repayments.payment, repayments.balance, loans.amount, repayments.created_at" +
		" FROM repayments INNER JOIN loans ON repayments.loan_id = loans.id WHERE repayments.loan_id = ?"

	qErr := db.Client.Select(installments, query, loanId)
	if qErr != nil {
		logger.Error(konstants.QUERY_ERR + qErr.Error())
		return nil, ericerrors.New500Error(konstants.MSG_500)
	}

	return &installments, nil
}
