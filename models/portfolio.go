package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Portfolio is a collection of Holdings and a target AssetAllocation.
type Portfolio struct {
	ID              uuid.UUID                          `json:"uuid,omitEmpty" db:"portfolio_uuid"`
	Name            string                             `json:"name" db:"name"`
	Description     string                             `json:"description" db:"description"`
	Holdings        []Holding                          `json:"holdings"`
	AssetAllocation map[*AssetCategory]decimal.Decimal `json:"-"`
}

// PortfolioDTO is the way to handle data transfer of Portfolio.
type PortfolioDTO struct {
	ID                             uuid.UUID                       `json:"uuid,omitEmpty" db:"portfolio_uuid"`
	Name                           string                          `json:"name" db:"name"`
	Description                    string                          `json:"description" db:"description"`
	PortfolioHoldingMappings       []PortfolioHoldingMapping       `json:"holdings"`
	PortfolioAssetCategoryMappings []PortfolioAssetCategoryMapping `json:"assetAllocation"`
}

// PortfolioHoldingMapping .
type PortfolioHoldingMapping struct {
	ID        uuid.UUID `json:"uuid,omitEmpty" db:"portfolio_holding_mapping_uuid"`
	Portfolio Portfolio `json:"portfolio" db:"portfolio"`
	Holding   Holding   `json:"holding" db:"holding"`
}

// PortfolioAssetCategoryMapping .
type PortfolioAssetCategoryMapping struct {
	ID            uuid.UUID       `json:"uuid,omitEmpty" db:"portfolio_asset_category_mapping_uuid"`
	Portfolio     Portfolio       `json:"portfolio" db:"portfolio"`
	AssetCategory AssetCategory   `json:"category" db:"asset_category"`
	Percentage    decimal.Decimal `json:"percentage" db:"percentage"`
}

// ToPortfolio converts a *PortfolioDTO to a Portfolio.
func (p *PortfolioDTO) ToPortfolio() *Portfolio {
	hs := make([]Holding, len(p.PortfolioHoldingMappings))
	for i, phm := range p.PortfolioHoldingMappings {
		hs[i] = phm.Holding
	}

	aas := make(map[*AssetCategory]decimal.Decimal)
	for _, pacm := range p.PortfolioAssetCategoryMappings {
		ac := &AssetCategory{
			ID: pacm.AssetCategory.ID,
		}
		aas[ac] = pacm.Percentage
	}

	return &Portfolio{
		ID:              p.ID,
		Name:            p.Name,
		Description:     p.Description,
		Holdings:        hs,
		AssetAllocation: aas,
	}
}

// ToPortfolioDTO converts a Portfolio to a PortfolioDTO.
// There is loss in the IDs of the PortfolioHoldingMappings and PortfolioAssetCategoryMappings.
func (p *Portfolio) ToPortfolioDTO() *PortfolioDTO {
	phms := make([]PortfolioHoldingMapping, len(p.Holdings))
	for i, h := range p.Holdings {
		phm := PortfolioHoldingMapping{
			Portfolio: *p,
			Holding:   h,
		}
		phms[i] = phm
	}

	var pacms []PortfolioAssetCategoryMapping
	for k, v := range p.AssetAllocation {
		pacm := PortfolioAssetCategoryMapping{
			Portfolio:     *p,
			AssetCategory: *k,
			Percentage:    v,
		}
		pacms = append(pacms, pacm)
	}

	return &PortfolioDTO{
		ID:                             p.ID,
		Name:                           p.Name,
		Description:                    p.Description,
		PortfolioHoldingMappings:       phms,
		PortfolioAssetCategoryMappings: pacms,
	}
}
