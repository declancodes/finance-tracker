package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

const (
	accountCategory = "account category"
	account         = "account"
	contribution    = "contribution"
)

// AccountController is the means for interacting with Account entities from an http router.
type AccountController struct{}

type accountsResponse struct {
	Accounts []*models.Account `json:"accounts"`
	Total    decimal.Decimal   `json:"total"`
}

type contributionsResponse struct {
	Contributions []*models.Contribution `json:"contributions"`
	Total         decimal.Decimal        `json:"total"`
}

var accountRepo = repositories.AccountRepository{}

// CreateAccountCategory creates an AccountCategory based on the r *http.Request Body.
func (c *AccountController) CreateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac *models.AccountCategory
		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestModel(w, accountCategory, err)
			return
		}

		ac.ID = uuid.New()
		acIDs, err := accountRepo.CreateAccountCategories(db, []*models.AccountCategory{ac})
		if err != nil {
			errorCreating(w, accountCategory, err)
			return
		}
		created(w, acIDs[0])
	}
}

// CreateAccount creates an Account based on the r *http.Request Body.
func (c *AccountController) CreateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a *models.Account
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestModel(w, account, err)
			return
		}

		a.ID = uuid.New()
		aIDs, err := accountRepo.CreateAccounts(db, []*models.Account{a})
		if err != nil {
			errorCreating(w, account, err)
			return
		}
		created(w, aIDs[0])
	}
}

// CreateContribution creates a Contribution based on the r *http.Request Body.
func (c *AccountController) CreateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c *models.Contribution
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestModel(w, contribution, err)
			return
		}

		c.ID = uuid.New()
		cIDs, err := accountRepo.CreateContributions(db, []*models.Contribution{c})
		if err != nil {
			errorCreating(w, contribution, err)
			return
		}
		created(w, cIDs[0])
	}
}

// GetAccountCategory gets an AccountCategory.
func (c *AccountController) GetAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		ac, err := accountRepo.GetAccountCategory(db, acID)
		if err != nil {
			errorExecuting(w, accountCategory, err)
			return
		}
		read(w, ac, accountCategory)
	}
}

// GetAccountCategories gets AccountCategory entities.
func (c *AccountController) GetAccountCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acs, err := accountRepo.GetAccountCategories(db, getFilters(r))
		if err != nil {
			errorExecuting(w, accountCategory, err)
			return
		}
		read(w, acs, accountCategory)
	}
}

// GetAccount gets an Account.
func (c *AccountController) GetAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		a, err := accountRepo.GetAccount(db, aID)
		if err != nil {
			errorExecuting(w, account, err)
			return
		}

		updateAccountValueFromHoldings(db, a)

		read(w, a, account)
	}
}

// GetAccounts gets Account entities.
func (c *AccountController) GetAccounts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		as, err := accountRepo.GetAccounts(db, getFilters(r))
		if err != nil {
			errorExecuting(w, account, err)
			return
		}

		var t decimal.Decimal
		for _, a := range as {
			updateAccountValueFromHoldings(db, a)
			t = t.Add(a.Amount)
		}

		resp := accountsResponse{
			Accounts: as,
			Total:    t,
		}

		read(w, resp, account)
	}
}

// GetContribution gets a Contribution.
func (c *AccountController) GetContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		c, err := accountRepo.GetContribution(db, cID)
		if err != nil {
			errorExecuting(w, contribution, err)
			return
		}
		read(w, c, contribution)
	}
}

// GetContributions gets Contribution entities.
func (c *AccountController) GetContributions(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := accountRepo.GetContributions(db, getFilters(r))
		if err != nil {
			errorExecuting(w, contribution, err)
			return
		}

		var t decimal.Decimal
		for _, c := range cs {
			t = t.Add(c.Amount)
		}

		resp := contributionsResponse{
			Contributions: cs,
			Total:         t,
		}

		read(w, resp, contribution)
	}
}

// UpdateAccountCategory updates an AccountCategory.
func (c *AccountController) UpdateAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var ac *models.AccountCategory
		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestModel(w, accountCategory, err)
			return
		}

		ac.ID = acID
		err = accountRepo.UpdateAccountCategory(db, ac)
		if err != nil {
			errorExecuting(w, accountCategory, err)
			return
		}
		updated(w, ac.ID)
	}
}

// UpdateAccount updates an Account.
func (c *AccountController) UpdateAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		aID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var a *models.Account
		err = json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			badRequestModel(w, account, err)
			return
		}

		a.ID = aID
		err = accountRepo.UpdateAccount(db, a)
		if err != nil {
			errorExecuting(w, account, err)
			return
		}
		updated(w, a.ID)
	}
}

// UpdateContribution updates a Contribution.
func (c *AccountController) UpdateContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var c *models.Contribution
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			badRequestModel(w, contribution, err)
			return
		}

		c.ID = cID
		err = accountRepo.UpdateContribution(db, c)
		if err != nil {
			errorExecuting(w, contribution, err)
			return
		}
		updated(w, c.ID)
	}
}

// DeleteAccountCategory deletes an AccountCategory.
func (c *AccountController) DeleteAccountCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, accountCategory, accountRepo.DeleteAccountCategory)
	}
}

// DeleteAccount deletes an Account.
func (c *AccountController) DeleteAccount(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, account, accountRepo.DeleteAccount)
	}
}

// DeleteContribution deletes a Contribution.
func (c *AccountController) DeleteContribution(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, contribution, accountRepo.DeleteContribution)
	}
}

func updateAccountValueFromHoldings(db *sqlx.DB, a *models.Account) {
	mValues := map[string]interface{}{
		"accounts": []string{a.Name},
	}
	hs, err := fundRepo.GetHoldings(db, mValues)
	if err != nil {
		log.Println(err)
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
