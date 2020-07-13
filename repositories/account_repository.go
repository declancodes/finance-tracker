package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AccountRepository is the means for interacting with Account storage.
type AccountRepository struct{}

// CreateAccountCategories creates AccountCategory entities in db.
func (r *AccountRepository) CreateAccountCategories(db *sqlx.DB, acs []*models.AccountCategory) ([]uuid.UUID, error) {
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

	IDs, err := createAndGetIDs(db, query, acs)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// CreateAccounts creates Account entities in db.
func (r *AccountRepository) CreateAccounts(db *sqlx.DB, as []*models.Account) ([]uuid.UUID, error) {
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

	IDs, err := createAndGetIDs(db, query, as)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// CreateContributions creates Contribution entities in db.
func (r *AccountRepository) CreateContributions(db *sqlx.DB, cs []*models.Contribution) ([]uuid.UUID, error) {
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

	IDs, err := createAndGetIDs(db, query, cs)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// GetAccountCategory retrieves the AccountCategory with acID from db.
func (r *AccountRepository) GetAccountCategory(db *sqlx.DB, acID uuid.UUID) (*models.AccountCategory, error) {
	mValues := map[string]interface{}{
		"account_category": acID.String(),
	}

	acs, err := r.GetAccountCategories(db, mValues)
	if err != nil {
		return nil, err
	}
	return acs[0], nil
}

// GetAccountCategories retrieves AccountCategory entities from db.
// Filters for AccountCategory retrieval are applied to the query based on the key-value pairs in mValues.
func (r *AccountRepository) GetAccountCategories(db *sqlx.DB, mValues map[string]interface{}) ([]*models.AccountCategory, error) {
	query := `
	SELECT
		account_category.account_category_uuid,
		account_category.name,
		account_category.description
	FROM account_category`

	mFilters := map[string]string{
		"account_category": "account_category.account_category_uuid = ",
	}

	q, args, err := getGetQueryAndValues(query, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var acs []*models.AccountCategory
	err = db.Select(&acs, q, args...)
	if err != nil {
		return nil, err
	}
	return acs, nil
}

// GetAccount retrieves the Account with aID from db.
func (r *AccountRepository) GetAccount(db *sqlx.DB, aID uuid.UUID) (*models.Account, error) {
	mValues := map[string]interface{}{
		"account": aID.String(),
	}

	as, err := r.GetAccounts(db, mValues)
	if err != nil {
		return nil, err
	}
	return as[0], nil
}

// GetAccounts retrieves Account entities from db.
// Filters for Account retrieval are applied to the query based on the key-value pairs in mValues.
func (r *AccountRepository) GetAccounts(db *sqlx.DB, mValues map[string]interface{}) ([]*models.Account, error) {
	query := `
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

	mFilters := map[string]string{
		"account":    "account.account_uuid = ",
		"categories": "account_category.name IN ",
	}

	q, args, err := getGetQueryAndValues(query, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var as []*models.Account
	err = db.Select(&as, q, args...)
	if err != nil {
		return nil, err
	}
	return as, nil
}

// GetContribution retrieves Contribution with cID from db.
func (r *AccountRepository) GetContribution(db *sqlx.DB, cID uuid.UUID) (*models.Contribution, error) {
	mValues := map[string]interface{}{
		"contribution": cID.String(),
	}

	cs, err := r.GetContributions(db, mValues)
	if err != nil {
		return nil, err
	}
	return cs[0], nil
}

// GetContributions retrieves Contribution entities from db.
// Filters for Contribution retrieval are applied to the query based on the key-value pairs in mValues.
func (r *AccountRepository) GetContributions(db *sqlx.DB, mValues map[string]interface{}) ([]*models.Contribution, error) {
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
		ON account.account_category_uuid = account_category.account_category_uuid`

	mFilters := map[string]string{
		"contribution": "contribution.contribution_uuid = ",
		"accounts":     "account.name IN ",
		"categories":   "account_category.name IN ",
		"start":        "contribution.date_made >= ",
		"end":          "contribution.date_made <= ",
	}

	q, args, err := getGetQueryAndValues(query, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	rows, err := db.Queryx(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cs []*models.Contribution
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

		cs = append(cs, &c)
	}
	return cs, nil
}

// UpdateAccountCategory updates an AccountCategory in db.
func (r *AccountRepository) UpdateAccountCategory(db *sqlx.DB, ac *models.AccountCategory) error {
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
func (r *AccountRepository) UpdateAccount(db *sqlx.DB, a *models.Account) error {
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
func (r *AccountRepository) UpdateContribution(db *sqlx.DB, c *models.Contribution) error {
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
func (r *AccountRepository) DeleteAccountCategory(db *sqlx.DB, acID uuid.UUID) error {
	query := `
	DELETE FROM account_category
	WHERE
		account_category_uuid = $1;`

	return deleteEntity(db, query, acID)
}

// DeleteAccount deletes an Account from db.
func (r *AccountRepository) DeleteAccount(db *sqlx.DB, aID uuid.UUID) error {
	query := `
	DELETE FROM account
	WHERE
		account_uuid = $1;`

	return deleteEntity(db, query, aID)
}

// DeleteContribution deletes a Contribution from db.
func (r *AccountRepository) DeleteContribution(db *sqlx.DB, cID uuid.UUID) error {
	query := `
	DELETE FROM contribution
	WHERE
		contribution_uuid = $1;`

	return deleteEntity(db, query, cID)
}
