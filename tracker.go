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
		w.Header().Set("Access-Control-Allow-Headers", "*")
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
	ec := controllers.ExpenseController{}
	fc := controllers.FundController{}
	pc := controllers.PortfolioController{}

	r := mux.NewRouter()

	r.HandleFunc("/accountcategories", ac.CreateAccountCategory(db)).Methods("POST")
	r.HandleFunc("/accountcategories", ac.GetAccountCategories(db)).Methods("GET")
	r.HandleFunc("/accountcategories/{id}", ac.GetAccountCategory(db)).Methods("GET")
	r.HandleFunc("/accountcategories/{id}", ac.UpdateAccountCategory(db)).Methods("PUT")
	r.HandleFunc("/accountcategories/{id}", ac.DeleteAccountCategory(db)).Methods("DELETE")

	r.HandleFunc("/accounts", ac.CreateAccount(db)).Methods("POST")
	r.HandleFunc("/accounts", ac.GetAccounts(db)).Methods("GET")
	r.HandleFunc("/accounts/{id}", ac.GetAccount(db)).Methods("GET")
	r.HandleFunc("/accounts/{id}", ac.UpdateAccount(db)).Methods("PUT")
	r.HandleFunc("/accounts/{id}", ac.DeleteAccount(db)).Methods("DELETE")

	r.HandleFunc("/contributions", ac.CreateContribution(db)).Methods("POST")
	r.HandleFunc("/contributions", ac.GetContributions(db)).Methods("GET")
	r.HandleFunc("/contributions/{id}", ac.GetContribution(db)).Methods("GET")
	r.HandleFunc("/contributions/{id}", ac.UpdateContribution(db)).Methods("PUT")
	r.HandleFunc("/contributions/{id}", ac.DeleteContribution(db)).Methods("DELETE")

	r.HandleFunc("/incomes", ac.CreateIncome(db)).Methods("POST")
	r.HandleFunc("/incomes", ac.GetIncomes(db)).Methods("GET")
	r.HandleFunc("/incomes/{id}", ac.GetIncome(db)).Methods("GET")
	r.HandleFunc("/incomes/{id}", ac.UpdateIncome(db)).Methods("PUT")
	r.HandleFunc("/incomes/{id}", ac.DeleteIncome(db)).Methods("DELETE")

	r.HandleFunc("/expensecategories", ec.CreateExpenseCategory(db)).Methods("POST")
	r.HandleFunc("/expensecategories", ec.GetExpenseCategories(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{id}", ec.GetExpenseCategory(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{id}", ec.UpdateExpenseCategory(db)).Methods("PUT")
	r.HandleFunc("/expensecategories/{id}", ec.DeleteExpenseCategory(db)).Methods("DELETE")

	r.HandleFunc("/expenses", ec.CreateExpense(db)).Methods("POST")
	r.HandleFunc("/expenses", ec.GetExpenses(db)).Methods("GET")
	r.HandleFunc("/expenses/{id}", ec.GetExpense(db)).Methods("GET")
	r.HandleFunc("/expenses/{id}", ec.UpdateExpense(db)).Methods("PUT")
	r.HandleFunc("/expenses/{id}", ec.DeleteExpense(db)).Methods("DELETE")

	r.HandleFunc("/assetcategories", fc.CreateAssetCategory(db)).Methods("POST")
	r.HandleFunc("/assetcategories", fc.GetAssetCategories(db)).Methods("GET")
	r.HandleFunc("/assetcategories/{id}", fc.GetAssetCategory(db)).Methods("GET")
	r.HandleFunc("/assetcategories/{id}", fc.UpdateAssetCategory(db)).Methods("PUT")
	r.HandleFunc("/assetcategories/{id}", fc.DeleteAssetCategory(db)).Methods("DELETE")

	r.HandleFunc("/funds", fc.CreateFund(db)).Methods("POST")
	r.HandleFunc("/funds", fc.GetFunds(db)).Methods("GET")
	r.HandleFunc("/funds", fc.UpdateFundSharePrices(db)).Methods("PUT")
	r.HandleFunc("/funds/{id}", fc.GetFund(db)).Methods("GET")
	r.HandleFunc("/funds/{id}", fc.UpdateFund(db)).Methods("PUT")
	r.HandleFunc("/funds/{id}", fc.DeleteFund(db)).Methods("DELETE")

	r.HandleFunc("/holdings", fc.CreateHolding(db)).Methods("POST")
	r.HandleFunc("/holdings", fc.GetHoldings(db)).Methods("GET")
	r.HandleFunc("/holdings/{id}", fc.GetHolding(db)).Methods("GET")
	r.HandleFunc("/holdings/{id}", fc.UpdateHolding(db)).Methods("PUT")
	r.HandleFunc("/holdings/{id}", fc.DeleteHolding(db)).Methods("DELETE")

	r.HandleFunc("/portfolios", pc.CreatePortfolio(db)).Methods("POST")
	r.HandleFunc("/portfolios", pc.GetPortfolios(db)).Methods("GET")
	r.HandleFunc("/portfolios/{id}", pc.GetPortfolio(db)).Methods("GET")
	r.HandleFunc("/portfolios/{id}", pc.UpdatePortfolio(db)).Methods("PUT")
	r.HandleFunc("/portfolios/{id}", pc.DeletePortfolio(db)).Methods("DELETE")

	r.HandleFunc("/portfolioholdingmappings", pc.CreatePortfolioHoldingMapping(db)).Methods("POST")
	r.HandleFunc("/portfolioholdingmappings", pc.GetPortfolioHoldingMappings(db)).Methods("GET")
	r.HandleFunc("/portfolioholdingmappings/{id}", pc.GetPortfolioHoldingMapping(db)).Methods("GET")
	r.HandleFunc("/portfolioholdingmappings/{id}", pc.UpdatePortfolioHoldingMapping(db)).Methods("PUT")
	r.HandleFunc("/portfolioholdingmappings/{id}", pc.DeletePortfolioHoldingMapping(db)).Methods("DELETE")

	r.HandleFunc("/portfolioassetcategorymappings", pc.CreatePortfolioAssetCategoryMapping(db)).Methods("POST")
	r.HandleFunc("/portfolioassetcategorymappings", pc.GetPortfolioAssetCategoryMappings(db)).Methods("GET")
	r.HandleFunc("/portfolioassetcategorymappings/{id}", pc.GetPortfolioAssetCategoryMapping(db)).Methods("GET")
	r.HandleFunc("/portfolioassetcategorymappings/{id}", pc.UpdatePortfolioAssetCategoryMapping(db)).Methods("PUT")
	r.HandleFunc("/portfolioassetcategorymappings/{id}", pc.DeletePortfolioAssetCategoryMapping(db)).Methods("DELETE")

	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(trailingSlashesMiddleware(r))))
}
