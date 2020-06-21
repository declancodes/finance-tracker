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
	AssetAllocation map[*AssetCategory]decimal.Decimal `json:"assetAllocation"`
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
	AssetCategory AssetCategory   `json:"assetCategory" db:"assetCategory"`
	Percentage    decimal.Decimal `json:"percentage" db:"percentage"`
}
