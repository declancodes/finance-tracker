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
	portfolio                     = "portfolio"
	portfolioHoldingMapping       = "portfolio holding mapping"
	portfolioAssetCategoryMapping = "portfolio asset category mapping"
)

// PortfolioController is the means for interacting with Portfolio entities from an http router.
type PortfolioController struct{}

var portfolioRepo = repositories.PortfolioRepository{}

// CreatePortfolio creates a Portfolio based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pDTO *models.PortfolioDTO
		err := json.NewDecoder(r.Body).Decode(&pDTO)
		if err != nil {
			badRequestModel(w, portfolio, err)
			return
		}

		p := pDTO.ToPortfolio()

		p.ID = uuid.New()
		pIDs, err := portfolioRepo.CreatePortfolios(db, []*models.Portfolio{p})
		if err != nil {
			errorCreating(w, portfolio, err)
			return
		}

		var phms []*models.PortfolioHoldingMapping
		for _, h := range p.Holdings {
			phmID := uuid.New()
			phms = append(phms, &models.PortfolioHoldingMapping{
				ID:        phmID,
				Portfolio: *p,
				Holding:   h,
			})
		}
		_, err = portfolioRepo.CreatePortfolioHoldingMappings(db, phms)
		if err != nil {
			errorCreating(w, portfolio, err)
			return
		}

		var pacms []*models.PortfolioAssetCategoryMapping
		for ac, per := range p.AssetAllocation {
			pacmID := uuid.New()
			pacms = append(pacms, &models.PortfolioAssetCategoryMapping{
				ID:            pacmID,
				Portfolio:     *p,
				AssetCategory: *ac,
				Percentage:    per,
			})

			log.Println(pacmID)
			log.Println(p.ID)
			log.Println(ac.ID)
			log.Println(per)
			log.Println()
		}
		_, err = portfolioRepo.CreatePortfolioAssetCategoryMappings(db, pacms)
		if err != nil {
			errorCreating(w, portfolio, err)
			return
		}
		created(w, pIDs[0])
	}
}

// CreatePortfolioHoldingMapping creates a PortfolioHoldingMapping based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var phm *models.PortfolioHoldingMapping
		err := json.NewDecoder(r.Body).Decode(&phm)
		if err != nil {
			badRequestModel(w, portfolioHoldingMapping, err)
			return
		}

		phm.ID = uuid.New()
		phmIDs, err := portfolioRepo.CreatePortfolioHoldingMappings(db, []*models.PortfolioHoldingMapping{phm})
		if err != nil {
			errorCreating(w, portfolioHoldingMapping, err)
			return
		}
		created(w, phmIDs[0])
	}
}

// CreatePortfolioAssetCategoryMapping creates a PortfolioAssetCategoryMapping based on the r *http.Request Body.
func (c *PortfolioController) CreatePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pacm *models.PortfolioAssetCategoryMapping
		err := json.NewDecoder(r.Body).Decode(&pacm)
		if err != nil {
			badRequestModel(w, portfolioAssetCategoryMapping, err)
			return
		}

		pacm.ID = uuid.New()
		pacmIDs, err := portfolioRepo.CreatePortfolioAssetCategoryMappings(db, []*models.PortfolioAssetCategoryMapping{pacm})
		if err != nil {
			errorCreating(w, portfolioAssetCategoryMapping, err)
			return
		}
		created(w, pacmIDs[0])
	}
}

// GetPortfolio gets a Portfolio.
func (c *PortfolioController) GetPortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		p, err := portfolioRepo.GetPortfolio(db, pID)
		if err != nil {
			errorExecuting(w, portfolio, err)
			return
		}
		pDTO := p.ToPortfolioDTO()
		read(w, pDTO, portfolio)
	}
}

// GetPortfolios gets Portfolio entities.
func (c *PortfolioController) GetPortfolios(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := portfolioRepo.GetPortfolios(db, getFilters(r))
		if err != nil {
			errorExecuting(w, portfolio, err)
			return
		}

		if len(ps) == 0 {
			errorExecuting(w, portfolio, err)
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
			errorExecuting(w, portfolio, err)
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
			errorExecuting(w, portfolio, err)
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

		pDTOs := make([]*models.PortfolioDTO, len(ps))
		for i, p := range ps {
			pDTOs[i] = p.ToPortfolioDTO()
		}

		read(w, pDTOs, portfolio)
	}
}

// GetPortfolioHoldingMapping gets a PortfolioHoldingMapping.
func (c *PortfolioController) GetPortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phmID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		phm, err := portfolioRepo.GetPortfolioHoldingMapping(db, phmID)
		if err != nil {
			errorExecuting(w, portfolioHoldingMapping, err)
			return
		}
		read(w, phm, portfolioHoldingMapping)
	}
}

// GetPortfolioHoldingMappings gets PortfolioHoldingMapping entities.
func (c *PortfolioController) GetPortfolioHoldingMappings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phms, err := portfolioRepo.GetPortfolioHoldingMappings(db, getFilters(r))
		if err != nil {
			errorExecuting(w, portfolioHoldingMapping, err)
			return
		}
		read(w, phms, portfolioHoldingMapping)
	}
}

// GetPortfolioAssetCategoryMapping gets a PortfolioAssetCategoryMapping.
func (c *PortfolioController) GetPortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pcamID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		pacm, err := portfolioRepo.GetPortfolioAssetCategoryMapping(db, pcamID)
		if err != nil {
			errorExecuting(w, portfolioAssetCategoryMapping, err)
			return
		}
		read(w, pacm, portfolioAssetCategoryMapping)
	}
}

// GetPortfolioAssetCategoryMappings gets PortfolioAssetCategoryMapping entities.
func (c *PortfolioController) GetPortfolioAssetCategoryMappings(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pacms, err := portfolioRepo.GetPortfolioAssetCategoryMappings(db, getFilters(r))
		if err != nil {
			errorExecuting(w, portfolioAssetCategoryMapping, err)
			return
		}
		read(w, pacms, portfolioAssetCategoryMapping)
	}
}

// UpdatePortfolio updates a Portfolio.
func (c *PortfolioController) UpdatePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var pDTO *models.PortfolioDTO
		err = json.NewDecoder(r.Body).Decode(&pDTO)
		if err != nil {
			badRequestModel(w, portfolio, err)
			return
		}

		p := pDTO.ToPortfolio()

		p.ID = pID
		err = portfolioRepo.UpdatePortfolio(db, p)
		if err != nil {
			errorExecuting(w, portfolio, err)
			return
		}
		updated(w, p.ID)
	}
}

// UpdatePortfolioHoldingMapping updates a PortfolioHoldingMapping.
func (c *PortfolioController) UpdatePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phmID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var phm *models.PortfolioHoldingMapping
		err = json.NewDecoder(r.Body).Decode(&phm)
		if err != nil {
			badRequestModel(w, portfolioHoldingMapping, err)
			return
		}

		phm.ID = phmID
		err = portfolioRepo.UpdatePortfolioHoldingMapping(db, phm)
		if err != nil {
			errorExecuting(w, portfolioHoldingMapping, err)
			return
		}
		updated(w, phm.ID)
	}
}

// UpdatePortfolioAssetCategoryMapping updates a PortfolioAssetCategoryMapping.
func (c *PortfolioController) UpdatePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pacmID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var pacm *models.PortfolioAssetCategoryMapping
		err = json.NewDecoder(r.Body).Decode(&pacm)
		if err != nil {
			badRequestModel(w, portfolioAssetCategoryMapping, err)
			return
		}

		pacm.ID = pacmID
		err = portfolioRepo.UpdatePortfolioAssetCategoryMapping(db, pacm)
		if err != nil {
			errorExecuting(w, portfolioAssetCategoryMapping, err)
			return
		}
		updated(w, pacm.ID)
	}
}

// DeletePortfolio deletes a Portfolio.
func (c *PortfolioController) DeletePortfolio(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, portfolio, portfolioRepo.DeletePortfolio)
	}
}

// DeletePortfolioHoldingMapping deletes a PortfolioHoldingMapping.
func (c *PortfolioController) DeletePortfolioHoldingMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, portfolioHoldingMapping, portfolioRepo.DeletePortfolioHoldingMapping)
	}
}

// DeletePortfolioAssetCategoryMapping deletes a PortfolioAssetCategoryMapping.
func (c *PortfolioController) DeletePortfolioAssetCategoryMapping(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, portfolioAssetCategoryMapping, portfolioRepo.DeletePortfolioAssetCategoryMapping)
	}
}
