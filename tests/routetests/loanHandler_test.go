package routetests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EricOgie/ope-be/app/conhandlers"
	"github.com/EricOgie/ope-be/domain/models"
	"github.com/EricOgie/ope-be/tests/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_get_user_loans_should_return_loans_with_status_200(t *testing.T) {
	// Define test controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Set up mock service for loan
	mockLoanServ := service.NewMockLoanServicePort(ctrl)

	dummyLoans := []models.QueryLoan{
		{Id: 3, Amount: 2400.365, Paid: 333.333, Package: "500 per month",
			Duration: "6 months", Status: "closed", CreatedAt: "02-02-21"},
		{Id: 3, Amount: 2400.365, Paid: 333.333, Package: "500 per month",
			Duration: "6 months", Status: "closed", CreatedAt: "02-02-21"},
		{Id: 3, Amount: 2400.365, Paid: 333.333, Package: "500 per month",
			Duration: "6 months", Status: "open", CreatedAt: "02-02-21"},
	}

	// Define
	mockLoanServ.EXPECT().GetLoans("").Return(dummyLoans, nil)

	loanHand := conhandlers.LoanHandler{mockLoanServ}

	router := mux.NewRouter()
	router.HandleFunc("/loans", loanHand.GetAllUserLoans)

	req, _ := http.NewRequest("GET", "/loans", nil)

	// ACT
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// ASSERT
	if rec.Code != http.StatusOK {
		t.Error("Failed status code while running Get loans")
	}
}
