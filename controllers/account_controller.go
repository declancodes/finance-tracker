package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
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

func badRequestContribution(w http.ResponseWriter, err error) {
	badRequestModel(w, "contribution", err)
}

func errorExecutingAccountCategory(w http.ResponseWriter, err error) {
	errorExecuting(w, "account category", err)
}

func errorExecutingAccount(w http.ResponseWriter, err error) {
	errorExecuting(w, "account", err)
}

func errorExecutingContribution(w http.ResponseWriter, err error) {
	errorExecuting(w, "contribution", err)
}

// CreateAccountCategory creates an AccountCategory based on the r *http.Request Body.
func (c *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac *models.AccountCategory
		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAccountCategory(w, err)
			return
		}

		ac.ID = uuid.New()
		acUUIDs, err := accountRepo.CreateAccountCategories(db, []*models.AccountCategory{ac})
		if err != nil {
			errorCreating(w, "account category", err)
			return
		}

		created(w, acUUIDs[0])
	}
}

// CreateAccount creates an Account based on the r *http.Request Body.
func (c *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a *models.Account
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestAccount(w, err)
			return
		}

		a.ID = uuid.New()
		aUUIDs, err := accountRepo.CreateAccounts(db, []*models.Account{a})
		if err != nil {
			errorCreating(w, "account", err)
			return
		}

		created(w, aUUIDs[0])
	}
}

// CreateContribution creates a Contribution based on the r *http.Request Body.
func (c *AccountController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c *models.Contribution
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestContribution(w, err)
			return
		}

		c.ID = uuid.New()
		cUUIDs, err := accountRepo.CreateContributions(db, []*models.Contribution{c})
		if err != nil {
			errorCreating(w, "contribution", err)
			return
		}

		created(w, cUUIDs[0])
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
		acs, err := accountRepo.GetAccountCategories(db, getFilters(r))
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

		updateAccountValueFromHoldings(db, a)

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(a)
		logError(err)
	}
}

// GetAccounts gets Account entities.
func (c *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		as, err := accountRepo.GetAccounts(db, getFilters(r))

		if err != nil {
			errorExecutingAccount(w, err)
			return
		}

		for _, a := range as {
			updateAccountValueFromHoldings(db, a)
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(as)
		logError(err)
	}
}

// GetContribution gets a Contribution.
func (c *AccountController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		c, err := accountRepo.GetContribution(db, cUUID)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(c)
		logError(err)
	}
}

// GetContributions gets Contribution entities.
func (c *AccountController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := accountRepo.GetContributions(db, getFilters(r))

		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(cs)
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

		var ac *models.AccountCategory
		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAccountCategory(w, err)
			return
		}

		ac.ID = acUUID
		err = accountRepo.UpdateAccountCategory(db, ac)
		if err != nil {
			errorExecutingAccountCategory(w, err)
			return
		}

		updated(w, ac.ID)
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

		var a *models.Account
		err = json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestAccount(w, err)
			return
		}

		a.ID = aUUID
		err = accountRepo.UpdateAccount(db, a)
		if err != nil {
			errorExecutingAccount(w, err)
			return
		}

		updated(w, a.ID)
	}
}

// UpdateContribution updates a Contribution.
func (c *AccountController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var c *models.Contribution
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestContribution(w, err)
			return
		}

		c.ID = cUUID
		err = accountRepo.UpdateContribution(db, c)
		if err != nil {
			errorExecutingContribution(w, err)
			return
		}

		updated(w, c.ID)
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

// DeleteContribution deletes a Contribution.
func (c *AccountController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "contribution", accountRepo.DeleteContribution)
	}
}

func updateAccountValueFromHoldings(db *sqlx.DB, a *models.Account) {
	mValues := map[string]interface{}{
		"accounts": []string{a.Name},
	}
	hs, err := fundRepo.GetHoldings(db, mValues)
	if err != nil {
		logError(err)
		return
	}

	if len(hs) > 0 {
		hTotal := decimal.Zero
		for _, h := range hs {
			hTotal = hTotal.Add(h.Value)
		}
		a.Amount = hTotal
	}
}
