package repositories

import (
	"fmt"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// HoldingRepository is the means for interacting with Holding storage.
type HoldingRepository struct{}

const (
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
		fund.name AS "fund.name",
		fund.ticker_symbol AS "fund.ticker_symbol",
		fund.share_price AS "fund.share_price",
		holding.shares
	FROM holding
	INNER JOIN account
		ON holding.account_uuid = account.account_uuid
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid
	INNER JOIN fund
		ON holding.fund_uuid = fund.fund_uuid`
)

// CreateHolding creates a Holding in db.
func (r *HoldingRepository) CreateHolding(db *sqlx.DB, h models.Holding) (uuid.UUID, error) {
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

	rows, err := db.NamedQuery(query, h)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&h.ID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return h.ID, nil
}

// GetHolding retrieves Holding with hUUID from db.
func (r *HoldingRepository) GetHolding(db *sqlx.DB, hUUID uuid.UUID) (h models.Holding, err error) {
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
func (r *HoldingRepository) GetHoldings(db *sqlx.DB, mValues map[string]interface{}) (hs []models.Holding, err error) {
	mFilters := map[string]string{
		"holding":  "holding.holding_uuid = ",
		"account":  "account.name = ",
		"category": "account_category.name = ",
		"fund":     "fund.ticker_symbol = ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return hs, err
	}

	query := fmt.Sprintf("%s %s", getHoldingsQuery, clauses)

	rows, err := db.Queryx(query, values...)
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
			&h.Fund.ID, &h.Fund.Name, &h.Fund.TickerSymbol, &h.Fund.SharePrice,
			&h.Shares)
		if err != nil {
			return hs, err
		}

		hs = append(hs, h)
	}
	return hs, nil
}

// UpdateHolding updates a Holding in db.
func (r *HoldingRepository) UpdateHolding(db *sqlx.DB, h models.Holding) error {
	query := `
	UPDATE holding
	SET
		account_uuid = :account.account_uuid,
		fund_uuid = :fund_uuid,
		shares = :shares
	WHERE
		holding_uuid = :holding_uuid;`

	res, err := db.NamedExec(query, h)

	_, err = getExecuted(res, err)
	return err
}

// DeleteHolding deletes a Holding from db.
func (r *HoldingRepository) DeleteHolding(db *sqlx.DB, hUUID uuid.UUID) error {
	query := `
	DELETE FROM holding
	WHERE
		holding_uuid = $1;`

	res, err := db.Exec(query, hUUID.String())

	_, err = getExecuted(res, err)
	return err
}
