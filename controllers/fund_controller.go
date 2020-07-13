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

const (
	assetCategory = "asset category"
	fund          = "fund"
	holding       = "holding"
)

// FundController is the means for interacting with Fund entities from an http router.
type FundController struct{}

var fundRepo = repositories.FundRepository{}

// CreateAssetCategory creates an AssetCategory based on the r *http.Request Body.
func (c *FundController) CreateAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ac *models.AssetCategory
		err := json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestModel(w, assetCategory, err)
			return
		}

		ac.ID = uuid.New()
		acIDs, err := fundRepo.CreateAssetCategories(db, []*models.AssetCategory{ac})
		if err != nil {
			errorCreating(w, assetCategory, err)
			return
		}
		created(w, acIDs[0])
	}
}

// CreateFund creates a Fund based on the r *http.Request Body.
func (c *FundController) CreateFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f *models.Fund
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			badRequestModel(w, fund, err)
			return
		}

		f.ID = uuid.New()
		f.TickerSymbol = strings.ToUpper(f.TickerSymbol)
		if f.SharePrice.Equal(decimal.Zero) {
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecuting(w, fund, err)
				return
			}
			f.SharePrice = sp
		}

		fIDs, err := fundRepo.CreateFunds(db, []*models.Fund{f})
		if err != nil {
			errorCreating(w, fund, err)
			return
		}
		created(w, fIDs[0])
	}
}

// CreateHolding creates a Holding based on the r *http.Request Body.
func (c *FundController) CreateHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var h *models.Holding
		err := json.NewDecoder(r.Body).Decode(&h)
		if err != nil {
			badRequestModel(w, holding, err)
			return
		}

		h.ID = uuid.New()
		hIDs, err := fundRepo.CreateHoldings(db, []*models.Holding{h})
		if err != nil {
			errorCreating(w, holding, err)
			return
		}
		created(w, hIDs[0])
	}
}

// GetAssetCategory gets an AssetCategory.
func (c *FundController) GetAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		ac, err := fundRepo.GetAssetCategory(db, acID)
		if err != nil {
			errorExecuting(w, assetCategory, err)
			return
		}
		read(w, ac, assetCategory)
	}
}

// GetAssetCategories gets AssetCategory entities.
func (c *FundController) GetAssetCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acs, err := fundRepo.GetAssetCategories(db, getFilters(r))
		if err != nil {
			errorExecuting(w, assetCategory, err)
			return
		}
		read(w, acs, assetCategory)
	}
}

// GetFund gets a Fund.
func (c *FundController) GetFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		f, err := fundRepo.GetFund(db, fID)
		if err != nil {
			errorExecuting(w, fund, err)
			return
		}
		read(w, f, fund)
	}
}

// GetFunds gets Fund entities.
func (c *FundController) GetFunds(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs, err := fundRepo.GetFunds(db, getFilters(r))
		if err != nil {
			errorExecuting(w, fund, err)
			return
		}
		read(w, fs, fund)
	}
}

// GetHolding gets a Holding.
func (c *FundController) GetHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		h, err := fundRepo.GetHolding(db, hID)
		if err != nil {
			errorExecuting(w, holding, err)
			return
		}
		read(w, h, holding)
	}
}

// GetHoldings gets Holding entities.
func (c *FundController) GetHoldings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hs, err := fundRepo.GetHoldings(db, getFilters(r))
		if err != nil {
			errorExecuting(w, holding, err)
			return
		}
		read(w, hs, holding)
	}
}

// UpdateAssetCategory updates an AssetCategory.
func (c *FundController) UpdateAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		acID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var ac *models.AssetCategory
		err = json.NewDecoder(r.Body).Decode(&ac)
		if err != nil {
			badRequestModel(w, assetCategory, err)
			return
		}

		ac.ID = acID
		err = fundRepo.UpdateAssetCategory(db, ac)
		if err != nil {
			errorExecuting(w, assetCategory, err)
			return
		}
		updated(w, ac.ID)
	}
}

// UpdateFund updates a Fund.
func (c *FundController) UpdateFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var f *models.Fund
		err = json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			badRequestModel(w, fund, err)
			return
		}

		f.ID = fID
		f.TickerSymbol = strings.ToUpper(f.TickerSymbol)
		if f.SharePrice.Equal(decimal.Zero) {
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecuting(w, fund, err)
				return
			}
			f.SharePrice = sp
		}

		err = fundRepo.UpdateFund(db, f)
		if err != nil {
			errorExecuting(w, fund, err)
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
			errorExecuting(w, fund, err)
			return
		}

		for _, f := range fs {
			sp, err := getSharePrice(f.TickerSymbol)
			if err != nil {
				errorExecuting(w, fund, err)
				return
			}
			f.SharePrice = sp

			err = fundRepo.UpdateFund(db, f)
			if err != nil {
				errorExecuting(w, fund, err)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateHolding updates a Holding.
func (c *FundController) UpdateHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var h *models.Holding
		err = json.NewDecoder(r.Body).Decode(&h)
		if err != nil {
			badRequestModel(w, holding, err)
			return
		}

		h.ID = hID
		err = fundRepo.UpdateHolding(db, h)
		if err != nil {
			errorExecuting(w, holding, err)
			return
		}
		updated(w, h.ID)
	}
}

// DeleteAssetCategory deletes an AssetCategory.
func (c *FundController) DeleteAssetCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, assetCategory, fundRepo.DeleteAssetCategory)
	}
}

// DeleteFund deletes a Fund.
func (c *FundController) DeleteFund(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, fund, fundRepo.DeleteFund)
	}
}

// DeleteHolding deletes a Holding.
func (c *FundController) DeleteHolding(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, holding, fundRepo.DeleteHolding)
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
