package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
