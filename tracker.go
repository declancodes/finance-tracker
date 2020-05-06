package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DeclanCodes/finance-tracker/controllers"
	"github.com/DeclanCodes/finance-tracker/driver"
	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: This is insecure! Specify allowed
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	})
}

func trailingSlashesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func main() {
	db := driver.DbConn()
	defer db.Close()

	ac := controllers.AccountController{}
	cc := controllers.ContributionController{}
	ec := controllers.ExpenseController{}
	hc := controllers.HoldingController{}

	r := mux.NewRouter()

	r.HandleFunc("/accountcategories", ac.CreateAccountCategory(db)).Methods("POST")
	r.HandleFunc("/accounts", ac.CreateAccount(db)).Methods("POST")
	r.HandleFunc("/accountcategories", ac.GetAccountCategories(db)).Methods("GET")
	r.HandleFunc("/accountcategories/{uuid}", ac.GetAccountCategory(db)).Methods("GET")
	r.HandleFunc("/accounts", ac.GetAccounts(db)).Methods("GET")
	r.HandleFunc("/accounts/{uuid}", ac.GetAccount(db)).Methods("GET")
	r.HandleFunc("/accountcategories/{uuid}", ac.UpdateAccountCategory(db)).Methods("PUT")
	r.HandleFunc("/accounts/{uuid}", ac.UpdateAccount(db)).Methods("PUT")
	r.HandleFunc("/accountcategories/{uuid}", ac.DeleteAccountCategory(db)).Methods("DELETE")
	r.HandleFunc("/accounts/{uuid}", ac.DeleteAccount(db)).Methods("DELETE")

	r.HandleFunc("/contributions", cc.CreateContribution(db)).Methods("POST")
	r.HandleFunc("/contributions", cc.GetContributions(db)).Methods("GET")
	r.HandleFunc("/contributions/{uuid}", cc.GetContribution(db)).Methods("GET")
	r.HandleFunc("/contributions/{uuid}", cc.UpdateContribution(db)).Methods("PUT")
	r.HandleFunc("/contributions/{uuid}", cc.DeleteContribution(db)).Methods("DELETE")

	r.HandleFunc("/expensecategories", ec.CreateExpenseCategory(db)).Methods("POST")
	r.HandleFunc("/expenses", ec.CreateExpense(db)).Methods("POST")
	r.HandleFunc("/expensecategories", ec.GetExpenseCategories(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{uuid}", ec.GetExpenseCategory(db)).Methods("GET")
	r.HandleFunc("/expenses", ec.GetExpenses(db)).Methods("GET")
	r.HandleFunc("/expenses/{uuid}", ec.GetExpense(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{uuid}", ec.UpdateExpenseCategory(db)).Methods("PUT")
	r.HandleFunc("/expenses/{uuid}", ec.UpdateExpense(db)).Methods("PUT")
	r.HandleFunc("/expensecategories/{uuid}", ec.DeleteExpenseCategory(db)).Methods("DELETE")
	r.HandleFunc("/expenses/{uuid}", ec.DeleteExpense(db)).Methods("DELETE")

	r.HandleFunc("/holdings", hc.CreateHolding(db)).Methods("POST")
	r.HandleFunc("/holdings", hc.GetHoldings(db)).Methods("GET")
	r.HandleFunc("/holdings/{uuid}", hc.GetHolding(db)).Methods("GET")
	r.HandleFunc("/holdings/{uuid}", hc.UpdateHolding(db)).Methods("PUT")
	r.HandleFunc("/holdings/{uuid}", hc.DeleteHolding(db)).Methods("DELETE")

	r.Use(corsMiddleware, loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8080", trailingSlashesMiddleware(r)))
}
