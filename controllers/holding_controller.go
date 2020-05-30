package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// HoldingController is the means for interacting with Holding entities from an http router.
type HoldingController struct{}

var holdingRepo = repositories.HoldingRepository{}

func badRequestHolding(w http.ResponseWriter, err error) {
	badRequestModel(w, "holding", err)
}

func errorExecutingHolding(w http.ResponseWriter, err error) {
	errorExecuting(w, "holding", err)
}

// CreateHolding creates a Holding based on the r *http.Request Body.
func (c *HoldingController) CreateHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var h models.Holding
		err := json.NewDecoder(r.Body).Decode(&h)
		if err != nil {
			badRequestHolding(w, err)
			return
		}

		h.ID, _ = uuid.NewUUID()
		hUUID, err := holdingRepo.CreateHolding(db, h)
		if err != nil {
			errorCreating(w, "holding", err)
			return
		}

		created(w, hUUID)
	}
}

// GetHolding gets a Holding.
func (c *HoldingController) GetHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		h, err := holdingRepo.GetHolding(db, hUUID)
		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		h.Value = h.Shares.Mul(h.Fund.SharePrice)

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(h)
		logError(err)
	}
}

// GetHoldings gets Holding entities.
func (c *HoldingController) GetHoldings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		accName := q.Get("account")
		catName := q.Get("category")
		fundSymbol := q.Get("fund")

		mValues := make(map[string]interface{})
		if accName != "" {
			mValues["account"] = accName
		}
		if catName != "" {
			mValues["category"] = catName
		}
		if fundSymbol != "" {
			mValues["fund"] = fundSymbol
		}

		hs, err := holdingRepo.GetHoldings(db, mValues)

		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		for _, h := range hs {
			h.Value = h.Shares.Mul(h.Fund.SharePrice)
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(hs)
		logError(err)
	}
}

// UpdateHolding updates a Holding.
func (c *HoldingController) UpdateHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var h models.Holding
		err = json.NewDecoder(r.Body).Decode(&h)
		if err != nil {
			badRequestHolding(w, err)
			return
		}

		h.ID = hUUID
		err = holdingRepo.UpdateHolding(db, h)
		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		updated(w, h.ID)
	}
}

// DeleteHolding deletes a Holding.
func (c *HoldingController) DeleteHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "holding", holdingRepo.DeleteHolding)
	}
}
