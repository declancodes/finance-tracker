package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

// FundController is the means for interacting with Fund entities from an http router.
type FundController struct{}

var fundRepo = repositories.FundRepository{}

func badRequestFund(w http.ResponseWriter, err error) {
	badRequestModel(w, "fund", err)
}

func errorExecutingFund(w http.ResponseWriter, err error) {
	errorExecuting(w, "fund", err)
}

// CreateFund creates a Fund based on the r *http.Request Body.
func (c *FundController) CreateFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Fund
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			badRequestFund(w, err)
			return
		}

		f.ID, _ = uuid.NewUUID()
		f.TickerSymbol = strings.ToUpper(f.TickerSymbol)
		if f.SharePrice.Equal(decimal.Zero) {
			f.SharePrice = getSharePrice(f.TickerSymbol)
		}

		fUUID, err := fundRepo.CreateFund(db, f)
		if err != nil {
			errorCreating(w, "fund", err)
			return
		}

		created(w, fUUID)
	}
}

// GetFund gets a Fund.
func (c *FundController) GetFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		f, err := fundRepo.GetFund(db, fUUID)
		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(f)
		logError(err)
	}
}

// GetFunds gets Fund entities.
func (c *FundController) GetFunds(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs, err := fundRepo.GetFunds(db)

		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(fs)
		logError(err)
	}
}

// UpdateFund updates a Fund.
func (c *FundController) UpdateFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var f models.Fund
		err = json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			badRequestFund(w, err)
			return
		}

		f.ID = fUUID
		f.TickerSymbol = strings.ToUpper(f.TickerSymbol)
		if f.SharePrice.Equal(decimal.Zero) {
			f.SharePrice = getSharePrice(f.TickerSymbol)
		}

		err = fundRepo.UpdateFund(db, f)
		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		updated(w, f.ID)
	}
}

// DeleteFund deletes a Fund.
func (c *FundController) DeleteFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "fund", fundRepo.DeleteFund)
	}
}

func getSharePrice(s string) decimal.Decimal {
	u := &url.URL{
		Scheme:   "https",
		Host:     os.Getenv("IEX_HOST"),
		Path:     fmt.Sprintf("v1/stock/%s/previous", s),
		RawQuery: fmt.Sprintf("token=%s", os.Getenv("IEX_KEY")),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	logError(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	logError(err)

	var pp models.PreviousPrice
	err = json.NewDecoder(resp.Body).Decode(&pp)
	logError(err)

	return pp.Close
}
