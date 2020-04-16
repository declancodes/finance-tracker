package repositories

import (
	"log"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AccountRepository is the means for interacting with Account storage.
type AccountRepository struct{}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// CreateAccountCategory creates an AccountCategory in db.
func (r *AccountRepository) CreateAccountCategory(db *sqlx.DB, ac models.AccountCategory) uuid.UUID {
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
	logError(err)

	for rows.Next() {
		err = rows.Scan(&ac.AccountCategoryUUID)
		logError(err)
	}

	return ac.AccountCategoryUUID
}

// CreateAccount creates an Account in db.
func (r *AccountRepository) CreateAccount(db *sqlx.DB, a models.Account) uuid.UUID {
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
	logError(err)

	for rows.Next() {
		err = rows.Scan(&a.AccountUUID)
		logError(err)
	}

	return a.AccountUUID
}

// GetAccountCategory retrieves an AccountCategory from db.
func (r *AccountRepository) GetAccountCategory(db *sqlx.DB, acUUID uuid.UUID) (ac models.AccountCategory) {
	query := `
	SELECT
		account_category_uuid,
		name,
		description
	FROM account_category
	WHERE
		account_category_uuid = $1;`

	err := db.Get(&ac, query, acUUID.String())
	logError(err)

	return ac
}

// GetAccountCategories retrieves AccountCategorys from db.
func (r *AccountRepository) GetAccountCategories(db *sqlx.DB) (acs []models.AccountCategory) {
	query := `
	SELECT
		account_category_uuid,
		name,
		description
	FROM account_category;`

	err := db.Select(&acs, query)
	logError(err)

	return acs
}

// GetAccount retrieves an Account from db.
func (r *AccountRepository) GetAccount(db *sqlx.DB, aUUID uuid.UUID) (a models.Account) {
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
		ON account.account_category_uuid = account_category.account_category_uuid
	WHERE
		account.account_uuid = $1;`

	err := db.Get(&a, query, aUUID.String())
	logError(err)

	return a
}

// GetAccounts retrieves Accounts from db.
func (r *AccountRepository) GetAccounts(db *sqlx.DB) (as []models.Account) {
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
		ON account.account_category_uuid = account_category.account_category_uuid;`

	err := db.Select(&as, query)
	logError(err)

	return as
}

// UpdateAccountCategory updates an AccountCategory in db.
func (r *AccountRepository) UpdateAccountCategory(db *sqlx.DB, ac models.AccountCategory) {
	query := `
	UPDATE account_category
	SET
		name = :name,
		description = :description
	WHERE
		account_category_uuid = :account_category_uuid;`

	_, err := db.NamedExec(query, ac)
	logError(err)
}

// UpdateAccount updates an Account in db.
func (r *AccountRepository) UpdateAccount(db *sqlx.DB, a models.Account) {
	query := `
	UPDATE account
	SET
		account_category_uuid = :account_category.account_category_uuid,
		name = :name,
		description = :description,
		amount = :amount
	WHERE
		account_uuid = :account_uuid;`

	_, err := db.NamedExec(query, a)
	logError(err)
}

// DeleteAccountCategory deletes an AccountCategory from db.
func (r *AccountRepository) DeleteAccountCategory(db *sqlx.DB, acUUID uuid.UUID) {
	query := `
	DELETE FROM account_category
	WHERE
		account_category_uuid = $1;`

	_, err := db.Exec(query, acUUID.String())
	logError(err)
}

// DeleteAccount deletes an Account from db.
func (r *AccountRepository) DeleteAccount(db *sqlx.DB, aUUID uuid.UUID) {
	query := `
	DELETE FROM account
	WHERE
		account_uuid = $1;`

	_, err := db.Exec(query, aUUID.String())
	logError(err)
}
