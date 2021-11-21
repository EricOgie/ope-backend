package repositories

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/ericerrors"
	"github.com/EricOgie/ope-be/konstants"
	"github.com/EricOgie/ope-be/logger"
)

//
func getLoan(loanId int, db LoanRepo) (*models.Loan, *error) {
	query := "SELECT * FROM loans WHER id = ?"

	var loan models.Loan
	err := db.Client.Get(loan, query, loanId)

	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return nil, &err
	}

	return &loan, nil
}

//
func updateLoan(db LoanRepo, l *models.Loan, pay float64) (string, *ericerrors.EricError) {
	loanBal := l.Amount - l.Paid
	var error error
	var status string
	if loanBal < pay {
		query := "UPDATE loans SET paid = paid + ? WHERE id = ?"
		_, e := db.Client.Exec(query, pay, l.Id)
		error = e
		status = "open"

	} else {
		query := "UPDATE loans SET paid = paid + ?, status = ? WHERE id = ?"
		_, e := db.Client.Exec(query, pay, "closed", l.Id)
		error = e
		if e != nil {
			status = "open"
		} else {
			status = "closed"
		}

	}

	ericErr := ericerrors.EricError{Code: 500, Message: error.Error()}
	return status, &ericErr
}

//
func getCorrectPaymnet(loan *models.Loan, pay models.LoanPayment) float64 {
	// Check id User is sending in more than is required to close th loan
	var correactPay float64
	loanBal := loan.Amount - loan.Paid

	if loanBal < pay.Payment { // User is sending more than is reqired
		correactPay = loanBal
	} else {
		correactPay = pay.Payment
	}

	return correactPay
}

//
//

func minusFromWallet(db LoanRepo, userId int, pay float64) *ericerrors.EricError {
	query := "UPDATE wallet SET amount = amount - ? WHERE user_id = ?"
	_, err := db.Client.Exec(query, pay, userId)
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		ericErr := ericerrors.NewError(http.StatusInternalServerError, konstants.MSG_500)
		return ericErr
	}
	return nil
}

func addToWallet(db LoanRepo, userId int, amount float64) *ericerrors.EricError {
	query := "UPDATE wallet SET amount = amount + ? WHERE user_id = ?"
	_, err := db.Client.Exec(query, amount, userId)
	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		ericErr := ericerrors.NewError(http.StatusInternalServerError, konstants.MSG_500)
		return ericErr
	}
	return nil
}

//
func CheckWallet(db LoanRepo, amount float64, userId int) bool {
	var walletFund float64
	query := "SELECT amount FROM wallet WHERE user_id = ?"
	err := db.Client.Get(&walletFund, query, userId)

	if err != nil {
		logger.Error(konstants.QUERY_ERR + err.Error())
		return false
	}

	return walletFund > amount
}

//
func Check60PercentMark(db LoanRepo, amount float64, userId int) bool {
	var potfolioPosition float64
	query := "SELECT SUM(equity_value) FROM stocks WHERE user_id = ?"
	qErr := db.Client.Get(&potfolioPosition, query, userId)

	if qErr != nil {
		logger.Error(konstants.QUERY_ERR + qErr.Error())
		return false
	}

	logger.Info("Position: " + strconv.Itoa(int(potfolioPosition)))
	return (0.6 * potfolioPosition) >= amount
}

func checkOpenLoans(db LoanRepo, userId int) bool {
	var amount float64
	query := "SELECT SUM(amount) FROM loans WHERE status = ? AND user_id = ?"
	err := db.Client.Get(&amount, query, "open", userId)
	if amount <= 0.000000 {
		logger.Error("NO OPEN LOAN err = " + fmt.Sprintf("%#v", err))
		return false
	}

	logger.Info("")
	return true
}
