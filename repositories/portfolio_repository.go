package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// PortfolioRepository is the means for interacting with Portfolio storage.
type PortfolioRepository struct{}

const (
	getPortfoliosQuery = `
	SELECT
		portfolio.portfolio_uuid,
		portfolio.name,
		portfolio.description
	FROM portfolio`

	getPortfolioHoldingMappingsQuery = `
	SELECT
		portfolio_holding_mapping.portfolio_holding_mapping_uuid,
		portfolio.portfolio_uuid AS "portfolio.portfolio_uuid",
		portfolio.name AS "portfolio.name",
		portfolio.description AS "portfolio.description",
		holding.holding_uuid AS "holding.holding_uuid",
		account.account_uuid AS "account.account_uuid",
		account_category.account_category_uuid AS "account_category.account_category_uuid",
		account_category.name AS "account_category.name",
		account_category.description AS "account_category.description",
		account.name AS "account.name",
		account.description AS "account.description",
		account.amount AS "account.amount",
		fund.fund_uuid AS "fund.fund_uuid",
		asset_category.asset_category_uuid AS "asset_category.asset_category_uuid",
		asset_category.name AS "asset_category.name",
		asset_category.description AS "asset_category.description",
		fund.name AS "fund.name",
		fund.ticker_symbol AS "fund.ticker_symbol",
		fund.share_price AS "fund.share_price",
		fund.expense_ratio AS "fund.expense_ratio",
		holding.shares
	FROM portfolio_holding_mapping
	INNER JOIN portfolio
		ON portfolio_holding_mapping.portfolio_uuid = portfolio.portfolio_uuid
	INNER JOIN holding
		ON portfolio_holding_mapping.holding_uuid = holding.holding_uuid
	INNER JOIN account
		ON holding.account_uuid = account.account_uuid
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid
	INNER JOIN fund
		ON holding.fund_uuid = fund.fund_uuid
	INNER JOIN asset_category
		ON fund.asset_category_uuid = asset_category.asset_category_uuid`

	getPortfolioAssetCategoryMappingsQuery = `
	SELECT
		portfolio_asset_category_mapping.portfolio_asset_category_mapping_uuid,
		portfolio.portfolio_uuid AS "portfolio.portfolio_uuid",
		portfolio.name AS "portfolio.name",
		portfolio.description AS "portfolio.description",
		asset_category.asset_category_uuid AS "asset_category.asset_category_uuid",
		asset_category.name AS "asset_category.name",
		asset_category.description AS "asset_category.description",
		portfolio_asset_category_mapping.percentage
	FROM portfolio_asset_category_mapping
	INNER JOIN portfolio
		ON portfolio_asset_category_mapping.portfolio_uuid = portfolio.portfolio_uuid
	INNER JOIN asset_category
		ON portfolio_asset_category_mapping.asset_category_uuid = asset_category.asset_category_uuid`
)

// CreatePortfolios creates Portfolio entities in db.
func (r *PortfolioRepository) CreatePortfolios(db *sqlx.DB, ps []*models.Portfolio) ([]uuid.UUID, error) {
	query := `
	INSERT INTO portfolio (
		portfolio_uuid,
		name,
		description
	)
	VALUES (
		:portfolio_uuid,
		:name,
		:description
	)
	RETURNING portfolio_uuid;`

	IDs, err := createAndGetUUIDs(db, query, ps)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// CreatePortfolioHoldingMappings creates PortfolioHoldingMapping entities in db.
func (r *PortfolioRepository) CreatePortfolioHoldingMappings(db *sqlx.DB, phms []*models.PortfolioHoldingMapping) ([]uuid.UUID, error) {
	query := `
	INSERT INTO portfolio_holding_mapping (
		portfolio_holding_mapping_uuid,
		portfolio_uuid,
		holding_uuid
	)
	VALUES (
		:portfolio_holding_mapping_uuid,
		:portfolio_uuid,
		:holding_uuid
	)
	RETURNING portfolio_holding_mapping_uuid;`

	IDs, err := createAndGetUUIDs(db, query, phms)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// CreatePortfolioAssetCategoryMappings creates PortfolioAssetCategoryMapping entities in db.
func (r *PortfolioRepository) CreatePortfolioAssetCategoryMappings(db *sqlx.DB, pacms []*models.PortfolioAssetCategoryMapping) ([]uuid.UUID, error) {
	query := `
	INSERT INTO portfolio_asset_category_mapping (
		portfolio_asset_category_mapping_uuid,
		portfolio_uuid,
		asset_category_uuid,
		percentage
	)
	VALUES (
		:portfolio_asset_category_mapping_uuid,
		:portfolio_uuid,
		:asset_category_uuid,
		:percentage
	)
	RETURNING portfolio_asset_category_mapping_uuid;`

	IDs, err := createAndGetUUIDs(db, query, pacms)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// GetPortfolio retrieves Portfolio with pUUID from db.
func (r *PortfolioRepository) GetPortfolio(db *sqlx.DB, pUUID uuid.UUID) (*models.Portfolio, error) {
	mValues := map[string]interface{}{
		"portfolio": pUUID.String(),
	}

	ps, err := r.GetPortfolios(db, mValues)
	if err != nil {
		return nil, err
	}

	return ps[0], nil
}

// GetPortfolios gets Portfolios from db.
func (r *PortfolioRepository) GetPortfolios(db *sqlx.DB, mValues map[string]interface{}) ([]*models.Portfolio, error) {
	mFilters := map[string]string{
		"portfolio": "portfolio.portfolio_uuid = ",
	}

	q, args, err := getGetQueryAndValues(getPortfoliosQuery, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var ps []*models.Portfolio
	err = db.Select(&ps, q, args...)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// GetPortfolioHoldingMapping retrieves PortfolioHoldingMapping with phmUUID from db.
func (r *PortfolioRepository) GetPortfolioHoldingMapping(db *sqlx.DB, phmUUID uuid.UUID) (*models.PortfolioHoldingMapping, error) {
	mValues := map[string]interface{}{
		"mapping": phmUUID.String(),
	}

	phms, err := r.GetPortfolioHoldingMappings(db, mValues)
	if err != nil {
		return nil, err
	}

	return phms[0], nil
}

// GetPortfolioHoldingMappings gets PortfolioHoldingMappings from db.
func (r *PortfolioRepository) GetPortfolioHoldingMappings(db *sqlx.DB, mValues map[string]interface{}) ([]*models.PortfolioHoldingMapping, error) {
	mFilters := map[string]string{
		"mapping":    "portfolio_holding_mapping.portfolio_holding_mapping_uuid = ",
		"portfolios": "portfolio.portfolio_uuid IN ",
	}

	q, args, err := getGetQueryAndValues(getPortfolioHoldingMappingsQuery, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var phms []*models.PortfolioHoldingMapping
	err = db.Select(&phms, q, args...)
	if err != nil {
		return nil, err
	}

	return phms, nil
}

// GetPortfolioAssetCategoryMapping retrieves PortfolioAssetCategoryMapping with pacmUUID from db.
func (r *PortfolioRepository) GetPortfolioAssetCategoryMapping(db *sqlx.DB, pacmUUID uuid.UUID) (*models.PortfolioAssetCategoryMapping, error) {
	mValues := map[string]interface{}{
		"mapping": pacmUUID.String(),
	}

	pacms, err := r.GetPortfolioAssetCategoryMappings(db, mValues)
	if err != nil {
		return nil, err
	}

	return pacms[0], nil
}

// GetPortfolioAssetCategoryMappings gets PortfolioAssetCategoryMappings from db.
func (r *PortfolioRepository) GetPortfolioAssetCategoryMappings(db *sqlx.DB, mValues map[string]interface{}) ([]*models.PortfolioAssetCategoryMapping, error) {
	mFilters := map[string]string{
		"mapping":    "portfolio_asset_category_mapping.portfolio_asset_category_mapping_uuid = ",
		"portfolios": "portfolio.portfolio_uuid IN ",
	}

	q, args, err := getGetQueryAndValues(getPortfolioAssetCategoryMappingsQuery, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var pacms []*models.PortfolioAssetCategoryMapping
	err = db.Select(&pacms, q, args...)
	if err != nil {
		return nil, err
	}

	return pacms, nil
}

// UpdatePortfolio updates a Portfolio in db.
func (r *PortfolioRepository) UpdatePortfolio(db *sqlx.DB, p *models.Portfolio) error {
	query := `
	UPDATE portfolio
	SET
		name = :name,
		description = :description
	WHERE
		portfolio_uuid = :portfolio_uuid;`

	return updateEntity(db, query, p)
}

// UpdatePortfolioHoldingMapping updates a PortfolioHoldingMapping in db.
func (r *PortfolioRepository) UpdatePortfolioHoldingMapping(db *sqlx.DB, phm *models.PortfolioHoldingMapping) error {
	query := `
	UPDATE portfolio_holding_mapping
	SET
		portfolio_uuid = :portfolio_uuid,
		holding_uuid = :holding_uuid
	WHERE
		portfolio_holding_mapping_uuid = :portfolio_holding_mapping_uuid;`

	return updateEntity(db, query, phm)
}

// UpdatePortfolioAssetCategoryMapping updates a PortfolioAssetCategoryMapping in db.
func (r *PortfolioRepository) UpdatePortfolioAssetCategoryMapping(db *sqlx.DB, pacm *models.PortfolioAssetCategoryMapping) error {
	query := `
	UPDATE portfolio_asset_category_mapping
	SET
		portfolio_uuid = :portfolio_uuid,
		asset_category_uuid = :asset_category_uuid,
		percentage = :percentage
	WHERE
		portfolio_asset_category_mapping_uuid = :portfolio_asset_category_mapping_uuid;`

	return updateEntity(db, query, pacm)
}

// DeletePortfolio deletes a Portfolio from db.
func (r *PortfolioRepository) DeletePortfolio(db *sqlx.DB, pUUID uuid.UUID) error {
	query := `
	DELETE FROM portfolio
	WHERE
		portfolio_uuid = $1;`

	return deleteEntity(db, query, pUUID)
}

// DeletePortfolioHoldingMapping deletes a PortfolioHoldingMapping from db.
func (r *PortfolioRepository) DeletePortfolioHoldingMapping(db *sqlx.DB, phmUUID uuid.UUID) error {
	query := `
	DELETE FROM portfolio_holding_mapping
	WHERE
		portfolio_holding_mapping_uuid = $1;`

	return deleteEntity(db, query, phmUUID)
}

// DeletePortfolioAssetCategoryMapping deletes a PortfolioAssetCategoryMapping from db.
func (r *PortfolioRepository) DeletePortfolioAssetCategoryMapping(db *sqlx.DB, pacmUUID uuid.UUID) error {
	query := `
	DELETE FROM portfolio_asset_category_mapping
	WHERE
		portfolio_asset_category_mapping_uuid = $1;`

	return deleteEntity(db, query, pacmUUID)
}
