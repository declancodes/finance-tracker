package repositories

import (
	"fmt"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// FundRepository is the means for interacting with Fund storage.
type FundRepository struct{}

const (
	getAssetCategoriesQuery = `
	SELECT
		asset_category_uuid,
		name,
		description
	FROM asset_category`

	getFundsQuery = `
	SELECT
		fund.fund_uuid,
		asset_category.asset_category_uuid AS "asset_category.asset_category_uuid",
		asset_category.name AS "asset_category.name",
		asset_category.description AS "asset_category.description",
		fund.name,
		fund.ticker_symbol,
		fund.share_price,
		fund.expense_ratio
	FROM fund
	INNER JOIN asset_category
		ON fund.asset_category_uuid = asset_category.asset_category_uuid`

	getHoldingsQuery = `
	SELECT
		holding.holding_uuid,
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
	FROM holding
	INNER JOIN account
		ON holding.account_uuid = account.account_uuid
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid
	INNER JOIN fund
		ON holding.fund_uuid = fund.fund_uuid
	INNER JOIN asset_category
		ON fund.asset_category_uuid = asset_category.asset_category_uuid`
)

// CreateAssetCategory creates an AssetCategory in db.
func (r *FundRepository) CreateAssetCategory(db *sqlx.DB, ac models.AssetCategory) (uuid.UUID, error) {
	query := `
	INSERT INTO asset_category (
		asset_category_uuid,
		name,
		description
	)
	VALUES (
		:asset_category_uuid,
		:name,
		:description
	)
	RETURNING asset_category_uuid;`

	return createAndGetUUID(db, query, ac)
}

// CreateFund creates a Fund in db.
func (r *FundRepository) CreateFund(db *sqlx.DB, f models.Fund) (uuid.UUID, error) {
	query := `
	INSERT INTO fund (
		fund_uuid,
		asset_category_uuid,
		name,
		ticker_symbol,
		share_price,
		expense_ratio
	)
	VALUES (
		:fund_uuid,
		:asset_category.asset_category_uuid,
		:name,
		:ticker_symbol,
		:share_price,
		:expense_ratio
	)
	RETURNING fund_uuid;`

	return createAndGetUUID(db, query, f)
}

// CreateHolding creates a Holding in db.
func (r *FundRepository) CreateHolding(db *sqlx.DB, h models.Holding) (uuid.UUID, error) {
	query := `
	INSERT INTO holding (
		holding_uuid,
		account_uuid,
		fund_uuid,
		shares
	)
	VALUES (
		:holding_uuid,
		:account.account_uuid,
		:fund.fund_uuid,
		:shares
	)
	RETURNING holding_uuid;`

	return createAndGetUUID(db, query, h)
}

// GetAssetCategory retrieves the AssetCategory with acUUID from db.
func (r *FundRepository) GetAssetCategory(db *sqlx.DB, acUUID uuid.UUID) (ac models.AssetCategory, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		asset_category_uuid = $1;`, getAssetCategoriesQuery)

	err = db.Get(&ac, query, acUUID.String())
	return ac, err
}

// GetAssetCategories retrieves AssetCategory entities from db.
func (r *FundRepository) GetAssetCategories(db *sqlx.DB) (acs []models.AssetCategory, err error) {
	query := fmt.Sprintf(`%s;`, getAssetCategoriesQuery)

	err = db.Select(&acs, query)
	return acs, err
}

// GetFund retrieves the Fund with fUUID from db.
func (r *FundRepository) GetFund(db *sqlx.DB, fUUID uuid.UUID) (f models.Fund, err error) {
	mValues := map[string]interface{}{
		"fund": fUUID.String(),
	}

	fs, err := r.GetFunds(db, mValues)
	if err != nil {
		return f, err
	}
	if len(fs) > 1 {
		return f, fmt.Errorf("more than one Fund with ID: %v", fUUID)
	}

	return fs[0], nil
}

// GetFunds retrieves Fund entities from db.
// Filters for Fund retrieval are applied to the query based on the key-value pairs in mValues.
func (r *FundRepository) GetFunds(db *sqlx.DB, mValues map[string]interface{}) (fs []models.Fund, err error) {
	mFilters := map[string]string{
		"fund":     "fund.fund_uuid = ",
		"category": "asset_category.name IN ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return fs, err
	}

	query := fmt.Sprintf("%s %s", getFundsQuery, clauses)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	err = db.Select(&fs, q, args...)
	return fs, err
}

// GetHolding retrieves Holding with hUUID from db.
func (r *FundRepository) GetHolding(db *sqlx.DB, hUUID uuid.UUID) (h *models.Holding, err error) {
	mValues := map[string]interface{}{
		"holding": hUUID.String(),
	}

	hs, err := r.GetHoldings(db, mValues)
	if err != nil {
		return h, err
	}
	if len(hs) > 1 {
		return h, fmt.Errorf("more than one Holding with ID: %v", hUUID)
	}

	return hs[0], nil
}

// GetHoldings retrieves Holding entities from db.
// Filters for Holding retrieval are applied to the query based on the key-value pairs in mValues.
func (r *FundRepository) GetHoldings(db *sqlx.DB, mValues map[string]interface{}) (hs []*models.Holding, err error) {
	mFilters := map[string]string{
		"holding":    "holding.holding_uuid = ",
		"accounts":   "account.name IN ",
		"categories": "account_category.name IN ",
		"funds":      "fund.ticker_symbol IN ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return hs, err
	}

	query := fmt.Sprintf("%s %s", getHoldingsQuery, clauses)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	rows, err := db.Queryx(q, args...)
	if err != nil {
		return hs, err
	}
	defer rows.Close()

	for rows.Next() {
		var h models.Holding

		err = rows.Scan(&h.ID,
			&h.Account.ID,
			&h.Account.Category.ID, &h.Account.Category.Name, &h.Account.Category.Description,
			&h.Account.Name, &h.Account.Description, &h.Account.Amount,
			&h.Fund.ID,
			&h.Fund.Category.ID, &h.Fund.Category.Name, &h.Fund.Category.Description,
			&h.Fund.Name, &h.Fund.TickerSymbol, &h.Fund.SharePrice, &h.Fund.ExpenseRatio,
			&h.Shares)
		if err != nil {
			return hs, err
		}

		h.Value = h.Shares.Mul(h.Fund.SharePrice)
		h.EffectiveExpense = h.Value.Mul(h.Fund.ExpenseRatio)

		hs = append(hs, &h)
	}
	return hs, nil
}

// UpdateAssetCategory updates an AssetCategory in db.
func (r *FundRepository) UpdateAssetCategory(db *sqlx.DB, ac models.AssetCategory) error {
	query := `
	UPDATE asset_category
	SET
		name = :name,
		description = :description
	WHERE
		asset_category_uuid = :asset_category_uuid;`

	return updateEntity(db, query, ac)
}

// UpdateFund updates a Fund in db.
func (r *FundRepository) UpdateFund(db *sqlx.DB, f models.Fund) error {
	query := `
	UPDATE fund
	SET
		asset_category_uuid = :asset_category.asset_category_uuid,
		name = :name,
		ticker_symbol = :ticker_symbol,
		share_price = :share_price,
		expense_ratio = :expense_ratio
	WHERE
		fund_uuid = :fund_uuid;`

	return updateEntity(db, query, f)
}

// UpdateHolding updates a Holding in db.
func (r *FundRepository) UpdateHolding(db *sqlx.DB, h models.Holding) error {
	query := `
	UPDATE holding
	SET
		account_uuid = :account.account_uuid,
		fund_uuid = :fund.fund_uuid,
		shares = :shares
	WHERE
		holding_uuid = :holding_uuid;`

	return updateEntity(db, query, h)
}

// DeleteAssetCategory deletes an AssetCategory from db.
func (r *FundRepository) DeleteAssetCategory(db *sqlx.DB, acUUID uuid.UUID) error {
	query := `
	DELETE FROM asset_category
	WHERE
		asset_category_uuid = $1;`

	return deleteEntity(db, query, acUUID)
}

// DeleteFund deletes a Fund from db.
func (r *FundRepository) DeleteFund(db *sqlx.DB, fUUID uuid.UUID) error {
	query := `
	DELETE FROM fund
	WHERE
		fund_uuid = $1;`

	return deleteEntity(db, query, fUUID)
}

// DeleteHolding deletes a Holding from db.
func (r *FundRepository) DeleteHolding(db *sqlx.DB, hUUID uuid.UUID) error {
	query := `
	DELETE FROM holding
	WHERE
		holding_uuid = $1;`

	return deleteEntity(db, query, hUUID)
}
