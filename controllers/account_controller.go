package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// AccountController .
type AccountController struct{}

func getUUID(r *http.Request) uuid.UUID {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["uuid"])
	if err != nil {
		log.Fatal("invalid uuid")
	}
	return ID
}

// CreateAccountCategory .
func (accountController *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountCategory models.AccountCategory
		accountCategoryUUID, _ := uuid.NewUUID()

		json.NewDecoder(r.Body).Decode(&accountCategory)
		accountCategory.AccountCategoryUUID = accountCategoryUUID

		accountRepo := repositories.AccountRepository{}
		accountCategoryUUID = accountRepo.CreateAccountCategory(db, accountCategory)

		json.NewEncoder(w).Encode(accountCategoryUUID)
	}
}

// CreateAccount .
func (accountController *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account models.Account
		accountUUID, _ := uuid.NewUUID()

		json.NewDecoder(r.Body).Decode(&account)
		account.AccountUUID = accountUUID

		accountRepo := repositories.AccountRepository{}
		accountUUID = accountRepo.CreateAccount(db, account)

		json.NewEncoder(w).Encode(accountUUID)
	}
}

// GetAccountCategory .
func (accountController *AccountController) GetAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountCategory models.AccountCategory
		accountCategoryUUID := getUUID(r)

		accountRepo := repositories.AccountRepository{}
		accountCategory = accountRepo.GetAccountCategory(db, accountCategoryUUID)

		json.NewEncoder(w).Encode(accountCategory)
	}
}

// GetAccountCategories .
func (accountController *AccountController) GetAccountCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepo := repositories.AccountRepository{}
		accountCategories := accountRepo.GetAccountCategories(db)

		json.NewEncoder(w).Encode(accountCategories)
	}
}

// GetAccount .
func (accountController *AccountController) GetAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account models.Account
		accountUUID := getUUID(r)

		accountRepo := repositories.AccountRepository{}
		account = accountRepo.GetAccount(db, accountUUID)

		json.NewEncoder(w).Encode(account)
	}
}

// GetAccounts .
func (accountController *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountRepo := repositories.AccountRepository{}
		accounts := accountRepo.GetAccounts(db)

		json.NewEncoder(w).Encode(accounts)
	}
}

// UpdateAccountCategory .
func (accountController *AccountController) UpdateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountCategory models.AccountCategory
		accountCategoryUUID := getUUID(r)

		json.NewDecoder(r.Body).Decode(&accountCategory)
		accountCategory.AccountCategoryUUID = accountCategoryUUID

		accountRepo := repositories.AccountRepository{}
		accountRepo.UpdateAccountCategory(db, accountCategory)

		json.NewEncoder(w).Encode(accountCategory)
	}
}

// UpdateAccount .
func (accountController *AccountController) UpdateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account models.Account
		accountUUID := getUUID(r)

		json.NewDecoder(r.Body).Decode(&account)
		account.AccountUUID = accountUUID

		accountRepo := repositories.AccountRepository{}
		accountRepo.UpdateAccount(db, account)

		json.NewEncoder(w).Encode(account)
	}
}

// DeleteAccountCategory .
func (accountController *AccountController) DeleteAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCategoryUUID := getUUID(r)

		accountRepo := repositories.AccountRepository{}
		accountRepo.DeleteAccountCategory(db, accountCategoryUUID)
	}
}

// DeleteAccount .
func (accountController *AccountController) DeleteAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountUUID := getUUID(r)

		accountRepo := repositories.AccountRepository{}
		accountRepo.DeleteAccount(db, accountUUID)
	}
}
