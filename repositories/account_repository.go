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

	rows, err := db.NamedQuery(query, ac)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ac.AccountCategoryUUID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return ac.AccountCategoryUUID, nil
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

	rows, err := db.NamedQuery(query, a)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&a.AccountUUID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return a.AccountUUID, nil
}

// GetAccountCategory retrieves the AccountCategory with acUUID from db.
func (r *AccountRepository) GetAccountCategory(db *sqlx.DB, acUUID uuid.UUID) (ac models.AccountCategory, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		account_category_uuid = $1;`, getAccountsQuery)

	err = db.Get(&ac, query, acUUID.String())
	return ac, err
}

// GetAccountCategories retrieves AccountCategorys from db.
func (r *AccountRepository) GetAccountCategories(db *sqlx.DB) (acs []models.AccountCategory, err error) {
	query := fmt.Sprintf(`%s;`, getAccountCategoriesQuery)

	err = db.Select(&acs, query)
	return acs, err
}

// GetAccount retrieves the Account with aUUID from db.
func (r *AccountRepository) GetAccount(db *sqlx.DB, aUUID uuid.UUID) (a models.Account, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		account.account_uuid = $1;`, getAccountsQuery)

	err = db.Get(&a, query, aUUID.String())
	return a, err
}

// GetAccounts retrieves Accounts from db.
func (r *AccountRepository) GetAccounts(db *sqlx.DB) (as []models.Account, err error) {
	query := fmt.Sprintf(`%s;`, getAccountsQuery)

	err = db.Select(&as, query)
	return as, err
}

// GetAccountsByCategory retrieves Accounts with AccountCategory acName from db.
func (r *AccountRepository) GetAccountsByCategory(db *sqlx.DB, acName string) (as []models.Account, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		account_category.name = $1;`, getAccountsQuery)

	err = db.Select(&as, query, acName)
	return as, err
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

	res, err := db.NamedExec(query, ac)

	_, err = getExecuted(res, err)
	return err
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

	res, err := db.NamedExec(query, a)

	_, err = getExecuted(res, err)
	return err
}

// DeleteAccountCategory deletes an AccountCategory from db.
func (r *AccountRepository) DeleteAccountCategory(db *sqlx.DB, acUUID uuid.UUID) error {
	query := `
	DELETE FROM account_category
	WHERE
		account_category_uuid = $1;`

	res, err := db.Exec(query, acUUID.String())

	_, err = getExecuted(res, err)
	return err
}

// DeleteAccount deletes an Account from db.
func (r *AccountRepository) DeleteAccount(db *sqlx.DB, aUUID uuid.UUID) error {
	query := `
	DELETE FROM account
	WHERE
		account_uuid = $1;`

	res, err := db.Exec(query, aUUID.String())

	_, err = getExecuted(res, err)
	return err
}
