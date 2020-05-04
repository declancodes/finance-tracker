package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AccountController is the means for interacting with Account entities from an http router.
type AccountController struct{}

var accountRepo = repositories.AccountRepository{}

func badRequestAccountCategory(w http.ResponseWriter, err error) {
	badRequestModel(w, "account category", err)
}

func badRequestAccount(w http.ResponseWriter, err error) {
	badRequestModel(w, "account", err)
}

func errorExecutingAccountCategory(w http.ResponseWriter, err error) {
	errorExecuting(w, "account category", err)
}

func errorExecutingAccount(w http.ResponseWriter, err error) {
	errorExecuting(w, "account", err)
}

// CreateAccountCategory creates an AccountCategory based on the r *http.Request Body.
func (c *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac models.AccountCategory
		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAccountCategory(w, err)
			return
		}

		ac.AccountCategoryUUID, _ = uuid.NewUUID()
		acUUID, err := accountRepo.CreateAccountCategory(db, ac)
		if err != nil {
			errorCreating(w, "account category", err)
			return
		}

		created(w, acUUID)
	}
}

// CreateAccount creates an Account based on the r *http.Request Body.
func (c *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a models.Account
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestAccount(w, err)
			return
		}

		a.AccountUUID, _ = uuid.NewUUID()
		aUUID, err := accountRepo.CreateAccount(db, a)
		if err != nil {
			errorCreating(w, "account", err)
			return
		}

		created(w, aUUID)
	}
}

// GetAccountCategory gets an AccountCategory.
func (c *AccountController) GetAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		ac, err := accountRepo.GetAccountCategory(db, acUUID)
		if err != nil {
			errorExecutingAccountCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(ac)
		logError(err)
	}
}

// GetAccountCategories gets AccountCategory entities.
func (c *AccountController) GetAccountCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acs, err := accountRepo.GetAccountCategories(db)
		if err != nil {
			errorExecutingAccountCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(acs)
		logError(err)
	}
}

// GetAccount gets an Account.
func (c *AccountController) GetAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		a, err := accountRepo.GetAccount(db, aUUID)
		if err != nil {
			errorExecutingAccount(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(a)
		logError(err)
	}
}

// GetAccounts gets Account entities.
func (c *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		catName := q.Get("category")

		mValues := make(map[string]interface{})
		if catName != "" {
			mValues["category"] = catName
		}

		as, err := accountRepo.GetAccounts(db, mValues)

		if err != nil {
			errorExecutingAccount(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(as)
		logError(err)
	}
}

// UpdateAccountCategory updates an AccountCategory.
func (c *AccountController) UpdateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var ac models.AccountCategory
		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAccountCategory(w, err)
			return
		}

		ac.AccountCategoryUUID = acUUID
		err = accountRepo.UpdateAccountCategory(db, ac)
		if err != nil {
			errorExecutingAccountCategory(w, err)
			return
		}

		updated(w, ac.AccountCategoryUUID)
	}
}

// UpdateAccount updates an Account.
func (c *AccountController) UpdateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var a models.Account
		err = json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestAccount(w, err)
			return
		}

		a.AccountUUID = aUUID
		err = accountRepo.UpdateAccount(db, a)
		if err != nil {
			errorExecutingAccount(w, err)
			return
		}

		updated(w, a.AccountUUID)
	}
}

// DeleteAccountCategory deletes an AccountCategory.
func (c *AccountController) DeleteAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "account category", accountRepo.DeleteAccountCategory)
	}
}

// DeleteAccount deletes an Account.
func (c *AccountController) DeleteAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "account", accountRepo.DeleteAccount)
	}
}
