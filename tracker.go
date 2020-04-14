package main

import (
	"log"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/controllers"
	"github.com/DeclanCodes/finance-tracker/driver"
	"github.com/gorilla/mux"
)

func main() {
	db := driver.DbConn()
	defer db.Close()

	router := mux.NewRouter()

	accountController := controllers.AccountController{}
	contributionController := controllers.ContributionController{}
	expenseController := controllers.ExpenseController{}

	router.HandleFunc("/accountcategories", accountController.CreateAccountCategory(db)).Methods("POST")
	router.HandleFunc("/accounts", accountController.CreateAccount(db)).Methods("POST")
	router.HandleFunc("/accountcategories", accountController.GetAccountCategories(db)).Methods("GET")
	router.HandleFunc("/accountcategories/{uuid}", accountController.GetAccountCategory(db)).Methods("GET")
	router.HandleFunc("/accounts", accountController.GetAccounts(db)).Methods("GET")
	router.HandleFunc("/accounts/{uuid}", accountController.GetAccount(db)).Methods("GET")
	router.HandleFunc("/accountcategories/{uuid}", accountController.UpdateAccountCategory(db)).Methods("PUT")
	router.HandleFunc("/accounts/{uuid}", accountController.UpdateAccount(db)).Methods("PUT")
	router.HandleFunc("/accountcategories/{uuid}", accountController.DeleteAccountCategory(db)).Methods("DELETE")
	router.HandleFunc("/accounts/{uuid}", accountController.DeleteAccount(db)).Methods("DELETE")

	router.HandleFunc("/contributions", contributionController.CreateContribution(db)).Methods("POST")
	router.HandleFunc("/contributions", contributionController.GetContributions(db)).Methods("GET")
	router.HandleFunc("/contributions/{uuid}", contributionController.GetContribution(db)).Methods("GET")
	router.HandleFunc("/contributions/{uuid}", contributionController.UpdateContribution(db)).Methods("PUT")
	router.HandleFunc("/contributions/{uuid}", contributionController.DeleteContribution(db)).Methods("DELETE")

	router.HandleFunc("/expensecategories", expenseController.CreateExpenseCategory(db)).Methods("POST")
	router.HandleFunc("/expenses", expenseController.CreateExpense(db)).Methods("POST")
	router.HandleFunc("/expensecategories", expenseController.GetExpenseCategories(db)).Methods("GET")
	router.HandleFunc("/expensecategories/{uuid}", expenseController.GetExpenseCategory(db)).Methods("GET")
	router.HandleFunc("/expenses", expenseController.GetExpenses(db)).Methods("GET")
	router.HandleFunc("/expenses/{uuid}", expenseController.GetExpense(db)).Methods("GET")
	router.HandleFunc("/expensecategories/{uuid}", expenseController.UpdateExpenseCategory(db)).Methods("PUT")
	router.HandleFunc("/expenses/{uuid}", expenseController.UpdateExpense(db)).Methods("PUT")
	router.HandleFunc("/expensecategories/{uuid}", expenseController.DeleteExpenseCategory(db)).Methods("DELETE")
	router.HandleFunc("/expenses/{uuid}", expenseController.DeleteExpense(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
