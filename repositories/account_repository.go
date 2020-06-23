package repositories

import (
	"fmt"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AccountRepository is the means for interacting with Account storage.
type AccountRepository struct{}

const (
	getAccountCategoriesQuery = `
	SELECT
		account_category_uuid,
		name,
		description
	FROM account_category`

	getAccountsQuery = `
	SELECT
		account.account_uuid,
		account_category.account_category_uuid AS "account_category.account_category_uuid",
		account_category.name AS "account_category.name",
		account_category.description AS "account_category.description",
		account.name,
		account.description,
		account.amount
	FROM account
	INNER JOIN account_category
		ON account.account_category_uuid = account_category.account_category_uuid`

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

// CreateAccountCategory creates an AccountCategory in db.
func (r *AccountRepository) CreateAccountCategory(db *sqlx.DB, ac models.AccountCategory) (uuid.UUID, error) {
	query := `
	INSERT INTO account_category (
		account_category_uuid,
		name,
		description
	)
	VALUES (
		:account_category_uuid,
		:name,
		:description
	)
	RETURNING account_category_uuid;`

	return createAndGetUUID(db, query, ac)
}

// CreateAccount creates an Account in db.
func (r *AccountRepository) CreateAccount(db *sqlx.DB, a models.Account) (uuid.UUID, error) {
	query := `
	INSERT INTO account (
		account_uuid,
		account_category_uuid,
		name,
		description,
		amount
	)
	VALUES (
		:account_uuid,
		:account_category.account_category_uuid,
		:name,
		:description,
		:amount
	)
	RETURNING account_uuid;`

	return createAndGetUUID(db, query, a)
}

// CreateContribution creates a Contribution in db.
func (r *AccountRepository) CreateContribution(db *sqlx.DB, c models.Contribution) (uuid.UUID, error) {
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

	return createAndGetUUID(db, query, c)
}

// GetAccountCategory retrieves the AccountCategory with acUUID from db.
func (r *AccountRepository) GetAccountCategory(db *sqlx.DB, acUUID uuid.UUID) (ac models.AccountCategory, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		account_category_uuid = $1;`, getAccountCategoriesQuery)

	err = db.Get(&ac, query, acUUID.String())
	return ac, err
}

// GetAccountCategories retrieves AccountCategory entities from db.
func (r *AccountRepository) GetAccountCategories(db *sqlx.DB) (acs []models.AccountCategory, err error) {
	query := fmt.Sprintf(`%s;`, getAccountCategoriesQuery)

	err = db.Select(&acs, query)
	return acs, err
}

// GetAccount retrieves the Account with aUUID from db.
func (r *AccountRepository) GetAccount(db *sqlx.DB, aUUID uuid.UUID) (a *models.Account, err error) {
	mValues := map[string]interface{}{
		"account": aUUID.String(),
	}

	as, err := r.GetAccounts(db, mValues)
	if err != nil {
		return a, err
	}

	return as[0], nil
}

// GetAccounts retrieves Account entities from db.
// Filters for Account retrieval are applied to the query based on the key-value pairs in mValues.
func (r *AccountRepository) GetAccounts(db *sqlx.DB, mValues map[string]interface{}) (as []*models.Account, err error) {
	mFilters := map[string]string{
		"account":    "account.account_uuid = ",
		"categories": "account_category.name IN ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s %s", getAccountsQuery, clauses)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	err = db.Select(&as, q, args...)
	return as, err
}

// GetContribution retrieves Contribution with cUUID from db.
func (r *AccountRepository) GetContribution(db *sqlx.DB, cUUID uuid.UUID) (c models.Contribution, err error) {
	mValues := map[string]interface{}{
		"contribution": cUUID.String(),
	}

	cs, err := r.GetContributions(db, mValues)
	if err != nil {
		return c, err
	}

	return cs[0], nil
}

// GetContributions retrieves Contribution entities from db.
// Filters for Contribution retrieval are applied to the query based on the key-value pairs in mValues.
func (r *AccountRepository) GetContributions(db *sqlx.DB, mValues map[string]interface{}) (cs []models.Contribution, err error) {
	mFilters := map[string]string{
		"contribution": "contribution.contribution_uuid = ",
		"accounts":     "account.name IN ",
		"categories":   "account_category.name IN ",
		"start":        "contribution.date_made >= ",
		"end":          "contribution.date_made <= ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s %s", getContributionsQuery, clauses)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	rows, err := db.Queryx(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Contribution

		err = rows.Scan(&c.ID,
			&c.Account.ID,
			&c.Account.Category.ID, &c.Account.Category.Name, &c.Account.Category.Description,
			&c.Account.Name, &c.Account.Description, &c.Account.Amount,
			&c.Name, &c.Description, &c.Amount, &c.Date)
		if err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}
	return cs, nil
}

// UpdateAccountCategory updates an AccountCategory in db.
func (r *AccountRepository) UpdateAccountCategory(db *sqlx.DB, ac models.AccountCategory) error {
	query := `
	UPDATE account_category
	SET
		name = :name,
		description = :description
	WHERE
		account_category_uuid = :account_category_uuid;`

	return updateEntity(db, query, ac)
}

// UpdateAccount updates an Account in db.
func (r *AccountRepository) UpdateAccount(db *sqlx.DB, a models.Account) error {
	query := `
	UPDATE account
	SET
		account_category_uuid = :account_category.account_category_uuid,
		name = :name,
		description = :description,
		amount = :amount
	WHERE
		account_uuid = :account_uuid;`

	return updateEntity(db, query, a)
}

// UpdateContribution updates a Contribution in db.
func (r *AccountRepository) UpdateContribution(db *sqlx.DB, c models.Contribution) error {
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

	return updateEntity(db, query, c)
}

// DeleteAccountCategory deletes an AccountCategory from db.
func (r *AccountRepository) DeleteAccountCategory(db *sqlx.DB, acUUID uuid.UUID) error {
	query := `
	DELETE FROM account_category
	WHERE
		account_category_uuid = $1;`

	return deleteEntity(db, query, acUUID)
}

// DeleteAccount deletes an Account from db.
func (r *AccountRepository) DeleteAccount(db *sqlx.DB, aUUID uuid.UUID) error {
	query := `
	DELETE FROM account
	WHERE
		account_uuid = $1;`

	return deleteEntity(db, query, aUUID)
}

// DeleteContribution deletes a Contribution from db.
func (r *AccountRepository) DeleteContribution(db *sqlx.DB, cUUID uuid.UUID) error {
	query := `
	DELETE FROM contribution
	WHERE
		contribution_uuid = $1;`

	return deleteEntity(db, query, cUUID)
}
