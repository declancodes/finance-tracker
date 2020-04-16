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
		return uuid.Nil, err
	}
	return ID, nil
}

func writeHeaderForBadRequest(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))

	log.Println(err)
}

func writeHeaderForBadRequestUUID(w http.ResponseWriter, err error) {
	writeHeaderForBadRequest(w, "invalid uuid", err)
}

func writeHeaderForBadRequestModel(w http.ResponseWriter, model string, err error) {
	writeHeaderForBadRequest(w, "invalid "+model, err)
}

// CreateAccountCategory .
func (c *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac models.AccountCategory

		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			writeHeaderForBadRequestModel(w, "account category", err)
			return
		}

		ac.AccountCategoryUUID, _ = uuid.NewUUID()
		acUUID := accountRepo.CreateAccountCategory(db, ac)

		err = json.NewEncoder(w).Encode(acUUID)
		logError(err)
	}
}

// CreateAccount .
func (c *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a models.Account

		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			writeHeaderForBadRequestModel(w, "account", err)
			return
		}

		a.AccountUUID, _ = uuid.NewUUID()
		aUUID := accountRepo.CreateAccount(db, a)

		err = json.NewEncoder(w).Encode(aUUID)
		logError(err)
	}
}

// GetAccountCategory .
func (c *AccountController) GetAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		ac := accountRepo.GetAccountCategory(db, acUUID)

		err = json.NewEncoder(w).Encode(ac)
		logError(err)
	}
}

// GetAccountCategories .
func (c *AccountController) GetAccountCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acs := accountRepo.GetAccountCategories(db)

		err := json.NewEncoder(w).Encode(acs)
		logError(err)
	}
}

// GetAccount .
func (c *AccountController) GetAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		a := accountRepo.GetAccount(db, aUUID)

		err = json.NewEncoder(w).Encode(a)
		logError(err)
	}
}

// GetAccounts .
func (c *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		as := accountRepo.GetAccounts(db)

		err := json.NewEncoder(w).Encode(as)
		logError(err)
	}
}

// UpdateAccountCategory .
func (c *AccountController) UpdateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var ac models.AccountCategory

		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			writeHeaderForBadRequestModel(w, "account category", err)
			return
		}

		ac.AccountCategoryUUID = acUUID
		accountRepo.UpdateAccountCategory(db, ac)

		err = json.NewEncoder(w).Encode(ac)
		logError(err)
	}
}

// UpdateAccount .
func (c *AccountController) UpdateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var a models.Account

		err = json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			writeHeaderForBadRequestModel(w, "account", err)
			return
		}

		a.AccountUUID = aUUID
		accountRepo.UpdateAccount(db, a)

		err = json.NewEncoder(w).Encode(a)
		logError(err)
	}
}

// DeleteAccountCategory .
func (c *AccountController) DeleteAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		accountRepo.DeleteAccountCategory(db, acUUID)
	}
}

// DeleteAccount .
func (c *AccountController) DeleteAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		accountRepo.DeleteAccount(db, aUUID)
	}
}
