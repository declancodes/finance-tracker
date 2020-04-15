package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ContributionRepository is the means for interacting with Contribution storage.
type ContributionRepository struct{}

// CreateContribution creates a Contribution in db.
func (contributionRepo *ContributionRepository) CreateContribution(db *sqlx.DB, contribution models.Contribution) uuid.UUID {
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

	rows, err := db.NamedQuery(query, contribution)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&contribution.ContributionUUID)
		logError(err)
	}

	return contribution.ContributionUUID
}

// GetContribution retrieves a Contribution from db.
func (contributionRepo *ContributionRepository) GetContribution(db *sqlx.DB, contributionUUID uuid.UUID) (contribution models.Contribution) {
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

	err := db.QueryRowx(query, contributionUUID.String()).Scan(&contribution.ContributionUUID,
		&contribution.Account.AccountUUID,
		&contribution.Account.AccountCategory.AccountCategoryUUID, &contribution.Account.AccountCategory.Name, &contribution.Account.AccountCategory.Description,
		&contribution.Account.Name, &contribution.Account.Description, &contribution.Account.Amount,
		&contribution.Name, &contribution.Description, &contribution.Amount, &contribution.Date)
	logError(err)

	return contribution
}

// GetContributions retrieves Contributions from db.
func (contributionRepo *ContributionRepository) GetContributions(db *sqlx.DB) (contributions []models.Contribution) {
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
	logError(err)

	for rows.Next() {
		var contribution models.Contribution

		err = rows.Scan(&contribution.ContributionUUID,
			&contribution.Account.AccountUUID,
			&contribution.Account.AccountCategory.AccountCategoryUUID, &contribution.Account.AccountCategory.Name, &contribution.Account.AccountCategory.Description,
			&contribution.Account.Name, &contribution.Account.Description, &contribution.Account.Amount,
			&contribution.Name, &contribution.Description, &contribution.Amount, &contribution.Date)
		logError(err)

		contributions = append(contributions, contribution)
	}

	return contributions
}

// UpdateContribution updates a Contribution in db.
func (contributionRepo *ContributionRepository) UpdateContribution(db *sqlx.DB, contribution models.Contribution) {
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

	_, err := db.NamedExec(query, contribution)
	logError(err)
}

// DeleteContribution deletes a Contribution from db.
func (contributionRepo *ContributionRepository) DeleteContribution(db *sqlx.DB, contributionUUID uuid.UUID) {
	query := `
	DELETE FROM contribution
	WHERE
		contribution_uuid = $1;`

	_, err := db.Exec(query, contributionUUID.String())
	logError(err)
}
