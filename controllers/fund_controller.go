package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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

func badRequestAssetCategory(w http.ResponseWriter, err error) {
	badRequestModel(w, "asset category", err)
}

func badRequestFund(w http.ResponseWriter, err error) {
	badRequestModel(w, "fund", err)
}

func badRequestHolding(w http.ResponseWriter, err error) {
	badRequestModel(w, "holding", err)
}

func errorExecutingAssetCategory(w http.ResponseWriter, err error) {
	errorExecuting(w, "asset category", err)
}

func errorExecutingFund(w http.ResponseWriter, err error) {
	errorExecuting(w, "fund", err)
}

func errorExecutingHolding(w http.ResponseWriter, err error) {
	errorExecuting(w, "holding", err)
}

// CreateAssetCategory creates an AssetCategory based on the r *http.Request Body.
func (c *FundController) CreateAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac models.AssetCategory
		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAssetCategory(w, err)
			return
		}

		ac.ID, _ = uuid.NewUUID()
		acUUID, err := fundRepo.CreateAssetCategory(db, ac)
		if err != nil {
			errorCreating(w, "asset category", err)
			return
		}

		created(w, acUUID)
	}
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
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecutingFund(w, err)
				return
			}
			f.SharePrice = sp
		}

		fUUID, err := fundRepo.CreateFund(db, f)
		if err != nil {
			errorCreating(w, "fund", err)
			return
		}

		created(w, fUUID)
	}
}

// CreateHolding creates a Holding based on the r *http.Request Body.
func (c *FundController) CreateHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var h models.Holding
		err := json.NewDecoder(r.Body).Decode(&h)
		if err != nil {
			badRequestHolding(w, err)
			return
		}

		h.ID, _ = uuid.NewUUID()
		hUUID, err := fundRepo.CreateHolding(db, h)
		if err != nil {
			errorCreating(w, "holding", err)
			return
		}

		created(w, hUUID)
	}
}

// GetAssetCategory gets an AssetCategory.
func (c *FundController) GetAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		ac, err := fundRepo.GetAssetCategory(db, acUUID)
		if err != nil {
			errorExecutingAssetCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(ac)
		logError(err)
	}
}

// GetAssetCategories gets AssetCategory entities.
func (c *FundController) GetAssetCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acs, err := fundRepo.GetAssetCategories(db)
		if err != nil {
			errorExecutingAssetCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(acs)
		logError(err)
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
		fs, err := fundRepo.GetFunds(db, getFilters(r))

		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(fs)
		logError(err)
	}
}

// GetHolding gets a Holding.
func (c *FundController) GetHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		h, err := fundRepo.GetHolding(db, hUUID)
		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(h)
		logError(err)
	}
}

// GetHoldings gets Holding entities.
func (c *FundController) GetHoldings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hs, err := fundRepo.GetHoldings(db, getFilters(r))

		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(hs)
		logError(err)
	}
}

// UpdateAssetCategory updates an AssetCategory.
func (c *FundController) UpdateAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var ac models.AssetCategory
		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestAssetCategory(w, err)
			return
		}

		ac.ID = acUUID
		err = fundRepo.UpdateAssetCategory(db, ac)
		if err != nil {
			errorExecutingAssetCategory(w, err)
			return
		}

		updated(w, ac.ID)
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
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecutingFund(w, err)
				return
			}
			f.SharePrice = sp
		}

		err = fundRepo.UpdateFund(db, f)
		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		updated(w, f.ID)
	}
}

// UpdateFundSharePrices updates SharePrices for all Funds matching query params filter(s).
func (c *FundController) UpdateFundSharePrices(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs, err := fundRepo.GetFunds(db, getFilters(r))
		if err != nil {
			errorExecutingFund(w, err)
			return
		}

		for _, f := range fs {
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecutingFund(w, err)
				return
			}
			f.SharePrice = sp

			err = fundRepo.UpdateFund(db, *f)
			if err != nil {
				errorExecutingFund(w, err)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdateHolding updates a Holding.
func (c *FundController) UpdateHolding(db *sqlx.DB) http.HandlerFunc {
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
		err = fundRepo.UpdateHolding(db, h)
		if err != nil {
			errorExecutingHolding(w, err)
			return
		}

		updated(w, h.ID)
	}
}

// DeleteAssetCategory deletes an AssetCategory.
func (c *FundController) DeleteAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "asset category", fundRepo.DeleteAssetCategory)
	}
}

// DeleteFund deletes a Fund.
func (c *FundController) DeleteFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "fund", fundRepo.DeleteFund)
	}
}

// DeleteHolding deletes a Holding.
func (c *FundController) DeleteHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "holding", fundRepo.DeleteHolding)
	}
}

func getSharePrice(s string) (decimal.Decimal, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     os.Getenv("IEX_HOST"),
		Path:     fmt.Sprintf("v1/stock/%s/previous", s),
		RawQuery: fmt.Sprintf("token=%s", os.Getenv("IEX_KEY")),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println(err)
		return decimal.Zero, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return decimal.Zero, err
	}

	var pp models.PreviousPrice
	err = json.NewDecoder(resp.Body).Decode(&pp)
	if err != nil {
		log.Println(err)
		return decimal.Zero, err
	}

	return pp.Close, nil
}
