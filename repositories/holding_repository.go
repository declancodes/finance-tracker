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
		holding.name,
		holding.ticker_symbol,
		holding.shares
	FROM holding
	INNER JOIN account
		ON holding.account_uuid = account.account_uuid
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid`
)

// CreateHolding creates a Holding in db.
func (r *HoldingRepository) CreateHolding(db *sqlx.DB, h models.Holding) (uuid.UUID, error) {
	query := `
	INSERT INTO holding (
		holding_uuid,
		account_uuid,
		name,
		ticker_symbol,
		shares
	)
	VALUES (
		:holding_uuid,
		:account.account_uuid,
		:name,
		:ticker_symbol,
		:shares
	)
	RETURNING holding_uuid;`

	rows, err := db.NamedQuery(query, h)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&h.HoldingUUID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return h.HoldingUUID, nil
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

		err = rows.Scan(&h.HoldingUUID,
			&h.Account.AccountUUID,
			&h.Account.AccountCategory.AccountCategoryUUID, &h.Account.AccountCategory.Name, &h.Account.AccountCategory.Description,
			&h.Account.Name, &h.Account.Description, &h.Account.Amount,
			&h.Name, &h.TickerSymbol, &h.Shares)
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
		name = :name,
		ticker_symbol = :ticker_symbol,
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
