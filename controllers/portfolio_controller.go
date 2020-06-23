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

// PortfolioController is the means for interacting with Portfolio entities from an http router.
type PortfolioController struct{}

var portfolioRepo = repositories.PortfolioRepository{}

func badRequestPortfolio(w http.ResponseWriter, err error) {
	badRequestModel(w, "portfolio", err)
}

func badRequestPortfolioHoldingMapping(w http.ResponseWriter, err error) {
	badRequestModel(w, "portfolio holding mapping", err)
}

func badRequestPortfolioAssetCategoryMapping(w http.ResponseWriter, err error) {
	badRequestModel(w, "portfolio asset category mapping", err)
}

func errorExecutingPortfolio(w http.ResponseWriter, err error) {
	errorExecuting(w, "portfolio", err)
}

func errorExecutingPortfolioHoldingMapping(w http.ResponseWriter, err error) {
	errorExecuting(w, "portfolio holding mapping", err)
}

func errorExecutingPortfolioAssetCategoryMapping(w http.ResponseWriter, err error) {
	errorExecuting(w, "portfolio asset category mapping", err)
}

// CreatePortfolio creates a Portfolio based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Portfolio
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			badRequestPortfolio(w, err)
			return
		}

		p.ID, _ = uuid.NewUUID()
		pUUIDs, err := portfolioRepo.CreatePortfolios(db, []*models.Portfolio{&p})
		if err != nil {
			errorCreating(w, "portfolio", err)
			return
		}

		var phms []*models.PortfolioHoldingMapping
		for _, h := range p.Holdings {
			phmID, _ := uuid.NewUUID()
			phms = append(phms, &models.PortfolioHoldingMapping{
				ID:        phmID,
				Portfolio: p,
				Holding:   h,
			})
		}
		_, err = portfolioRepo.CreatePortfolioHoldingMappings(db, phms)
		if err != nil {
			errorCreating(w, "portfolio", err)
			return
		}

		var pacms []*models.PortfolioAssetCategoryMapping
		for ac, per := range p.AssetAllocation {
			pacmID, _ := uuid.NewUUID()
			pacms = append(pacms, &models.PortfolioAssetCategoryMapping{
				ID:            pacmID,
				Portfolio:     p,
				AssetCategory: *ac,
				Percentage:    per,
			})
		}
		_, err = portfolioRepo.CreatePortfolioAssetCategoryMappings(db, pacms)
		if err != nil {
			errorCreating(w, "portfolio", err)
			return
		}

		created(w, pUUIDs[0])
	}
}

// CreatePortfolioHoldingMapping creates a PortfolioHoldingMapping based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var phm models.PortfolioHoldingMapping
		err := json.NewDecoder(r.Body).Decode(&phm)
		if err != nil {
			badRequestPortfolioHoldingMapping(w, err)
			return
		}

		phm.ID, _ = uuid.NewUUID()
		phmUUIDs, err := portfolioRepo.CreatePortfolioHoldingMappings(db, []*models.PortfolioHoldingMapping{&phm})
		if err != nil {
			errorCreating(w, "portfolio holding mapping", err)
			return
		}

		created(w, phmUUIDs[0])
	}
}

// CreatePortfolioAssetCategoryMapping creates a PortfolioAssetCategoryMapping based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pacm models.PortfolioAssetCategoryMapping
		err := json.NewDecoder(r.Body).Decode(&pacm)
		if err != nil {
			badRequestPortfolioAssetCategoryMapping(w, err)
			return
		}

		pacm.ID, _ = uuid.NewUUID()
		pacmUUIDs, err := portfolioRepo.CreatePortfolioAssetCategoryMappings(db, []*models.PortfolioAssetCategoryMapping{&pacm})
		if err != nil {
			errorCreating(w, "portfolio asset category mapping", err)
			return
		}

		created(w, pacmUUIDs[0])
	}
}

// GetPortfolio gets a Portfolio.
func (c *PortfolioController) GetPortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		p, err := portfolioRepo.GetPortfolio(db, pUUID)
		if err != nil {
			errorExecutingPortfolio(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(p)
		logError(err)
	}
}

// GetPortfolios gets Portfolio entities.
func (c *PortfolioController) GetPortfolios(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := portfolioRepo.GetPortfolios(db, getFilters(r))
		if err != nil {
			errorExecutingPortfolio(w, err)
			return
		}

		pIDs := make([]uuid.UUID, len(ps))
		for i, p := range ps {
			pIDs[i] = p.ID
		}
		mValues := map[string]interface{}{
			"portfolios": pIDs,
		}

		phms, err := portfolioRepo.GetPortfolioHoldingMappings(db, mValues)
		if err != nil {
			errorExecutingPortfolio(w, err)
			return
		}
		phmsMap := make(map[string][]models.Holding)
		for _, phm := range phms {
			pID := phm.Portfolio.ID.String()
			if _, ok := phmsMap[pID]; ok {
				phmsMap[pID] = append(phmsMap[pID], phm.Holding)
			} else {
				phmsMap[pID] = []models.Holding{phm.Holding}
			}
		}

		pacms, err := portfolioRepo.GetPortfolioAssetCategoryMappings(db, mValues)
		if err != nil {
			errorExecutingPortfolio(w, err)
			return
		}
		pacmsMap := make(map[string]map[*models.AssetCategory]decimal.Decimal)
		for _, pacm := range pacms {
			pID := pacm.Portfolio.ID.String()
			if _, ok := pacmsMap[pID]; ok {
				pacmsMap[pID][&pacm.AssetCategory] = pacm.Percentage
			} else {
				pacmsMap[pID] = map[*models.AssetCategory]decimal.Decimal{
					&pacm.AssetCategory: pacm.Percentage,
				}
			}
		}

		for _, p := range ps {
			pID := p.ID.String()
			if hs, ok := phmsMap[pID]; ok {
				p.Holdings = hs
			}
			if ms, ok := pacmsMap[pID]; ok {
				p.AssetAllocation = ms
			}
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(ps)
		logError(err)
	}
}

// GetPortfolioHoldingMapping gets a PortfolioHoldingMapping.
func (c *PortfolioController) GetPortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phmUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		f, err := portfolioRepo.GetPortfolioHoldingMapping(db, phmUUID)
		if err != nil {
			errorExecutingPortfolioHoldingMapping(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(f)
		logError(err)
	}
}

// GetPortfolioHoldingMappings gets PortfolioHoldingMapping entities.
func (c *PortfolioController) GetPortfolioHoldingMappings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phms, err := portfolioRepo.GetPortfolioHoldingMappings(db, getFilters(r))

		if err != nil {
			errorExecutingPortfolioHoldingMapping(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(phms)
		logError(err)
	}
}

// GetPortfolioAssetCategoryMapping gets a PortfolioAssetCategoryMapping.
func (c *PortfolioController) GetPortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pcamUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		pcam, err := portfolioRepo.GetPortfolioAssetCategoryMapping(db, pcamUUID)
		if err != nil {
			errorExecutingPortfolioAssetCategoryMapping(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(pcam)
		logError(err)
	}
}

// GetPortfolioAssetCategoryMappings gets PortfolioAssetCategoryMapping entities.
func (c *PortfolioController) GetPortfolioAssetCategoryMappings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hs, err := portfolioRepo.GetPortfolioAssetCategoryMappings(db, getFilters(r))

		if err != nil {
			errorExecutingPortfolioAssetCategoryMapping(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(hs)
		logError(err)
	}
}

// UpdatePortfolio updates a Portfolio.
func (c *PortfolioController) UpdatePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var p models.Portfolio
		err = json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			badRequestPortfolio(w, err)
			return
		}

		p.ID = pUUID
		err = portfolioRepo.UpdatePortfolio(db, p)
		if err != nil {
			errorExecutingPortfolio(w, err)
			return
		}

		updated(w, p.ID)
	}
}

// UpdatePortfolioHoldingMapping updates a PortfolioHoldingMapping.
func (c *PortfolioController) UpdatePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phmUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var phm models.PortfolioHoldingMapping
		err = json.NewDecoder(r.Body).Decode(&phm)
		if err != nil {
			badRequestPortfolioHoldingMapping(w, err)
			return
		}

		phm.ID = phmUUID
		err = portfolioRepo.UpdatePortfolioHoldingMapping(db, phm)
		if err != nil {
			errorExecutingPortfolioHoldingMapping(w, err)
			return
		}

		updated(w, phm.ID)
	}
}

// UpdatePortfolioAssetCategoryMapping updates a PortfolioAssetCategoryMapping.
func (c *PortfolioController) UpdatePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pacmUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var pacm models.PortfolioAssetCategoryMapping
		err = json.NewDecoder(r.Body).Decode(&pacm)
		if err != nil {
			badRequestPortfolioAssetCategoryMapping(w, err)
			return
		}

		pacm.ID = pacmUUID
		err = portfolioRepo.UpdatePortfolioAssetCategoryMapping(db, pacm)
		if err != nil {
			errorExecutingPortfolioAssetCategoryMapping(w, err)
			return
		}

		updated(w, pacm.ID)
	}
}

// DeletePortfolio deletes a Portfolio.
func (c *PortfolioController) DeletePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "portfolio", portfolioRepo.DeletePortfolio)
	}
}

// DeletePortfolioHoldingMapping deletes a PortfolioHoldingMapping.
func (c *PortfolioController) DeletePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "portfolio holding mapping", portfolioRepo.DeletePortfolioHoldingMapping)
	}
}

// DeletePortfolioAssetCategoryMapping deletes a PortfolioAssetCategoryMapping.
func (c *PortfolioController) DeletePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "portfolio asset category mapping", portfolioRepo.DeletePortfolioAssetCategoryMapping)
	}
}
