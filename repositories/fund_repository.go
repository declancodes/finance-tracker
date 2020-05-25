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
	getFundsQuery = `
	SELECT
		fund_uuid,
		name,
		ticker_symbol,
		share_price
	FROM fund`
)

// CreateFund creates a Fund in db.
func (r *FundRepository) CreateFund(db *sqlx.DB, f models.Fund) (uuid.UUID, error) {
	query := `
	INSERT INTO fund (
		fund_uuid,
		name,
		ticker_symbol,
		share_price
	)
	VALUES (
		:fund_uuid,
		:name,
		:ticker_symbol,
		:share_price
	)
	RETURNING fund_uuid;`

	rows, err := db.NamedQuery(query, f)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&f.ID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return f.ID, nil
}

// GetFund retrieves the Fund with fUUID from db.
func (r *FundRepository) GetFund(db *sqlx.DB, fUUID uuid.UUID) (f models.Fund, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		fund.fund_uuid = $1;`, getFundsQuery)

	err = db.Get(&f, query, fUUID.String())
	return f, err
}

// GetFunds retrieves Fund entities from db.
func (r *FundRepository) GetFunds(db *sqlx.DB) (fs []models.Fund, err error) {
	query := fmt.Sprintf(`%s;`, getFundsQuery)

	err = db.Select(&fs, query)
	return fs, err
}

// UpdateFund updates a Fund in db.
func (r *FundRepository) UpdateFund(db *sqlx.DB, f models.Fund) error {
	query := `
	UPDATE fund
	SET
		name = :name,
		ticker_symbol = :ticker_symbol,
		share_price = :share_price
	WHERE
		fund_uuid = :fund_uuid;`

	res, err := db.NamedExec(query, f)

	_, err = getExecuted(res, err)
	return err
}

// DeleteFund deletes a Fund from db.
func (r *FundRepository) DeleteFund(db *sqlx.DB, fUUID uuid.UUID) error {
	query := `
	DELETE FROM fund
	WHERE
		fund_uuid = $1;`

	res, err := db.Exec(query, fUUID.String())

	_, err = getExecuted(res, err)
	return err
}
