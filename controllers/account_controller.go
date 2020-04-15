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

var accountRepo = repositories.AccountRepository{}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getUUID(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	ID, err := uuid.Parse(params["uuid"])
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}
	return ID, nil
}

// CreateAccountCategory .
func (accountController *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accountCategory models.AccountCategory
		err := json.NewDecoder(r.Body).Decode(&accountCategory)
		logError(err)

		accountCategory.AccountCategoryUUID, _ = uuid.NewUUID()
		accountCategoryUUID := accountRepo.CreateAccountCategory(db, accountCategory)

		err = json.NewEncoder(w).Encode(accountCategoryUUID)
		logError(err)
	}
}

// CreateAccount .
func (accountController *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account models.Account
		err := json.NewDecoder(r.Body).Decode(&account)
		logError(err)

		account.AccountUUID, _ = uuid.NewUUID()
		accountUUID := accountRepo.CreateAccount(db, account)

		err = json.NewEncoder(w).Encode(accountUUID)
		logError(err)
	}
}

// GetAccountCategory .
func (accountController *AccountController) GetAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCategoryUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		accountCategory := accountRepo.GetAccountCategory(db, accountCategoryUUID)

		err = json.NewEncoder(w).Encode(accountCategory)
		logError(err)
	}
}

// GetAccountCategories .
func (accountController *AccountController) GetAccountCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCategories := accountRepo.GetAccountCategories(db)

		err := json.NewEncoder(w).Encode(accountCategories)
		logError(err)
	}
}

// GetAccount .
func (accountController *AccountController) GetAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		account := accountRepo.GetAccount(db, accountUUID)

		err = json.NewEncoder(w).Encode(account)
		logError(err)
	}
}

// GetAccounts .
func (accountController *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accounts := accountRepo.GetAccounts(db)

		err := json.NewEncoder(w).Encode(accounts)
		logError(err)
	}
}

// UpdateAccountCategory .
func (accountController *AccountController) UpdateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCategoryUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		var accountCategory models.AccountCategory
		err = json.NewDecoder(r.Body).Decode(&accountCategory)
		logError(err)

		accountCategory.AccountCategoryUUID = accountCategoryUUID
		accountRepo.UpdateAccountCategory(db, accountCategory)

		err = json.NewEncoder(w).Encode(accountCategory)
		logError(err)
	}
}

// UpdateAccount .
func (accountController *AccountController) UpdateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		var account models.Account
		err = json.NewDecoder(r.Body).Decode(&account)
		logError(err)

		account.AccountUUID = accountUUID
		accountRepo.UpdateAccount(db, account)

		err = json.NewEncoder(w).Encode(account)
		logError(err)
	}
}

// DeleteAccountCategory .
func (accountController *AccountController) DeleteAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCategoryUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		accountRepo.DeleteAccountCategory(db, accountCategoryUUID)
	}
}

// DeleteAccount .
func (accountController *AccountController) DeleteAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountUUID, err := getUUID(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid uuid"))
			return
		}

		accountRepo.DeleteAccount(db, accountUUID)
	}
}
