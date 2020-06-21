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
	r.HandleFunc("/accountcategories/{uuid}", ac.GetAccountCategory(db)).Methods("GET")
	r.HandleFunc("/accountcategories/{uuid}", ac.UpdateAccountCategory(db)).Methods("PUT")
	r.HandleFunc("/accountcategories/{uuid}", ac.DeleteAccountCategory(db)).Methods("DELETE")

	r.HandleFunc("/accounts", ac.CreateAccount(db)).Methods("POST")
	r.HandleFunc("/accounts", ac.GetAccounts(db)).Methods("GET")
	r.HandleFunc("/accounts/{uuid}", ac.GetAccount(db)).Methods("GET")
	r.HandleFunc("/accounts/{uuid}", ac.UpdateAccount(db)).Methods("PUT")
	r.HandleFunc("/accounts/{uuid}", ac.DeleteAccount(db)).Methods("DELETE")

	r.HandleFunc("/contributions", ac.CreateContribution(db)).Methods("POST")
	r.HandleFunc("/contributions", ac.GetContributions(db)).Methods("GET")
	r.HandleFunc("/contributions/{uuid}", ac.GetContribution(db)).Methods("GET")
	r.HandleFunc("/contributions/{uuid}", ac.UpdateContribution(db)).Methods("PUT")
	r.HandleFunc("/contributions/{uuid}", ac.DeleteContribution(db)).Methods("DELETE")

	r.HandleFunc("/expensecategories", ec.CreateExpenseCategory(db)).Methods("POST")
	r.HandleFunc("/expensecategories", ec.GetExpenseCategories(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{uuid}", ec.GetExpenseCategory(db)).Methods("GET")
	r.HandleFunc("/expensecategories/{uuid}", ec.UpdateExpenseCategory(db)).Methods("PUT")
	r.HandleFunc("/expensecategories/{uuid}", ec.DeleteExpenseCategory(db)).Methods("DELETE")

	r.HandleFunc("/expenses", ec.CreateExpense(db)).Methods("POST")
	r.HandleFunc("/expenses", ec.GetExpenses(db)).Methods("GET")
	r.HandleFunc("/expenses/{uuid}", ec.GetExpense(db)).Methods("GET")
	r.HandleFunc("/expenses/{uuid}", ec.UpdateExpense(db)).Methods("PUT")
	r.HandleFunc("/expenses/{uuid}", ec.DeleteExpense(db)).Methods("DELETE")

	r.HandleFunc("/assetcategories", fc.CreateAssetCategory(db)).Methods("POST")
	r.HandleFunc("/assetcategories", fc.GetAssetCategories(db)).Methods("GET")
	r.HandleFunc("/assetcategories/{uuid}", fc.GetAssetCategory(db)).Methods("GET")
	r.HandleFunc("/assetcategories/{uuid}", fc.UpdateAssetCategory(db)).Methods("PUT")
	r.HandleFunc("/assetcategories/{uuid}", fc.DeleteAssetCategory(db)).Methods("DELETE")

	r.HandleFunc("/funds", fc.CreateFund(db)).Methods("POST")
	r.HandleFunc("/funds", fc.GetFunds(db)).Methods("GET")
	r.HandleFunc("/funds", fc.UpdateFundSharePrices(db)).Methods("PUT")
	r.HandleFunc("/funds/{uuid}", fc.GetFund(db)).Methods("GET")
	r.HandleFunc("/funds/{uuid}", fc.UpdateFund(db)).Methods("PUT")
	r.HandleFunc("/funds/{uuid}", fc.DeleteFund(db)).Methods("DELETE")

	r.HandleFunc("/holdings", fc.CreateHolding(db)).Methods("POST")
	r.HandleFunc("/holdings", fc.GetHoldings(db)).Methods("GET")
	r.HandleFunc("/holdings/{uuid}", fc.GetHolding(db)).Methods("GET")
	r.HandleFunc("/holdings/{uuid}", fc.UpdateHolding(db)).Methods("PUT")
	r.HandleFunc("/holdings/{uuid}", fc.DeleteHolding(db)).Methods("DELETE")

	r.HandleFunc("/portfolios", pc.CreatePortfolio(db)).Methods("POST")
	r.HandleFunc("/portfolios", pc.GetPortfolios(db)).Methods("GET")
	r.HandleFunc("/portfolios/{uuid}", pc.GetPortfolio(db)).Methods("GET")
	r.HandleFunc("/portfolios/{uuid}", pc.UpdatePortfolio(db)).Methods("PUT")
	r.HandleFunc("/portfolios/{uuid}", pc.DeletePortfolio(db)).Methods("DELETE")

	r.HandleFunc("/portfolioholdingmappings", pc.CreatePortfolioHoldingMapping(db)).Methods("POST")
	r.HandleFunc("/portfolioholdingmappings", pc.GetPortfolioHoldingMappings(db)).Methods("GET")
	r.HandleFunc("/portfolioholdingmappings/{uuid}", pc.GetPortfolioHoldingMapping(db)).Methods("GET")
	r.HandleFunc("/portfolioholdingmappings/{uuid}", pc.UpdatePortfolioHoldingMapping(db)).Methods("PUT")
	r.HandleFunc("/portfolioholdingmappings/{uuid}", pc.DeletePortfolioHoldingMapping(db)).Methods("DELETE")

	r.HandleFunc("/portfolioassetcategorymappings", pc.CreatePortfolioAssetCategoryMapping(db)).Methods("POST")
	r.HandleFunc("/portfolioassetcategorymappings", pc.GetPortfolioAssetCategoryMappings(db)).Methods("GET")
	r.HandleFunc("/portfolioassetcategorymappings/{uuid}", pc.GetPortfolioAssetCategoryMapping(db)).Methods("GET")
	r.HandleFunc("/portfolioassetcategorymappings/{uuid}", pc.UpdatePortfolioAssetCategoryMapping(db)).Methods("PUT")
	r.HandleFunc("/portfolioassetcategorymappings/{uuid}", pc.DeletePortfolioAssetCategoryMapping(db)).Methods("DELETE")

	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(trailingSlashesMiddleware(r))))
}
