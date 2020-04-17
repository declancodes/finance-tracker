package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ContributionRepository is the means for interacting with Contribution storage.
type ContributionRepository struct{}

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

// GetContribution retrieves a Contribution from db.
func (r *ContributionRepository) GetContribution(db *sqlx.DB, cUUID uuid.UUID) (c models.Contribution, err error) {
	query := `
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
		ON account.account_category_uuid = account_category.account_category_uuid
	WHERE
		contribution.contribution_uuid = $1;`

	err = db.QueryRowx(query, cUUID.String()).Scan(&c.ContributionUUID,
		&c.Account.AccountUUID,
		&c.Account.AccountCategory.AccountCategoryUUID, &c.Account.AccountCategory.Name, &c.Account.AccountCategory.Description,
		&c.Account.Name, &c.Account.Description, &c.Account.Amount,
		&c.Name, &c.Description, &c.Amount, &c.Date)
	return c, err
}

// GetContributions retrieves Contributions from db.
func (r *ContributionRepository) GetContributions(db *sqlx.DB) (cs []models.Contribution, err error) {
	query := `
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
		ON account.account_category_uuid = account_category.account_category_uuid;`

	rows, err := db.Queryx(query)
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
