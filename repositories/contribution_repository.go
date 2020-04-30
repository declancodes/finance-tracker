package repositories

import (
	"fmt"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ContributionRepository is the means for interacting with Contribution storage.
type ContributionRepository struct{}

const (
	getContributionsQuery = `
	SELECT
		contribution.contribution_uuid,
		account.account_uuid AS "account.account_uuid",
		account_category.account_category_uuid AS "account_category.account_category_uuid",
		account_category.name AS "account_category.name",
		account_category.description AS "account_category.description",
		account.name AS "account.name",
		account.description AS "account.description",
		account.amount AS "account.amount",
		contribution.name,
		contribution.description,
		contribution.amount,
		contribution.date_made
	FROM contribution
	INNER JOIN account
		ON contribution.account_uuid = account.account_uuid
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid`
)

// CreateContribution creates a Contribution in db.
func (r *ContributionRepository) CreateContribution(db *sqlx.DB, c models.Contribution) (uuid.UUID, error) {
	query := `
	INSERT INTO contribution (
		contribution_uuid,
		account_uuid,
		name,
		description,
		amount,
		date_made
	)
	VALUES (
		:contribution_uuid,
		:account.account_uuid,
		:name,
		:description,
		:amount,
		:date_made
	)
	RETURNING contribution_uuid;`

	rows, err := db.NamedQuery(query, c)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.ContributionUUID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return c.ContributionUUID, nil
}

// GetContribution retrieves Contribution with cUUID from db.
func (r *ContributionRepository) GetContribution(db *sqlx.DB, cUUID uuid.UUID) (c models.Contribution, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		contribution.contribution_uuid = $1;`, getContributionsQuery)

	err = db.QueryRowx(query, cUUID.String()).Scan(&c.ContributionUUID,
		&c.Account.AccountUUID,
		&c.Account.AccountCategory.AccountCategoryUUID, &c.Account.AccountCategory.Name, &c.Account.AccountCategory.Description,
		&c.Account.Name, &c.Account.Description, &c.Account.Amount,
		&c.Name, &c.Description, &c.Amount, &c.Date)
	return c, err
}

// GetContributions retrieves Contribution entities from db.
// Filters for Contribution retrieval are applied to the query based on the key-value pairs in mValues.
func (r *ContributionRepository) GetContributions(db *sqlx.DB, mValues map[string]interface{}) (cs []models.Contribution, err error) {
	mFilters := map[string]string{
		"account":  "account.name = ",
		"category": "account_category.name = ",
		"start":    "contribution.date_made >= ",
		"end":      "contribution.date_made <= ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return cs, err
	}

	query := fmt.Sprintf("%s %s", getContributionsQuery, clauses)

	return getContributions(db, query, values...)
}

// UpdateContribution updates a Contribution in db.
func (r *ContributionRepository) UpdateContribution(db *sqlx.DB, c models.Contribution) error {
	query := `
	UPDATE contribution
	SET
		account_uuid = :account.account_uuid,
		name = :name,
		description = :description,
		amount = :amount,
		date_made = :date_made
	WHERE
		contribution_uuid = :contribution_uuid;`

	res, err := db.NamedExec(query, c)

	_, err = getExecuted(res, err)
	return err
}

// DeleteContribution deletes a Contribution from db.
func (r *ContributionRepository) DeleteContribution(db *sqlx.DB, cUUID uuid.UUID) error {
	query := `
	DELETE FROM contribution
	WHERE
		contribution_uuid = $1;`

	res, err := db.Exec(query, cUUID.String())

	_, err = getExecuted(res, err)
	return err
}

func getContributions(db *sqlx.DB, query string, args ...interface{}) (cs []models.Contribution, err error) {
	rows, err := db.Queryx(query, args...)
	if err != nil {
		return cs, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Contribution

		err = rows.Scan(&c.ContributionUUID,
			&c.Account.AccountUUID,
			&c.Account.AccountCategory.AccountCategoryUUID, &c.Account.AccountCategory.Name, &c.Account.AccountCategory.Description,
			&c.Account.Name, &c.Account.Description, &c.Account.Amount,
			&c.Name, &c.Description, &c.Amount, &c.Date)
		if err != nil {
			return cs, err
		}

		cs = append(cs, c)
	}
	return cs, nil
}
