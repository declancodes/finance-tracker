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
		log.Fatalln(err)
	}
}

// CreateAccountCategory creates an AccountCategory in db.
func (accountRepo *AccountRepository) CreateAccountCategory(db *sqlx.DB, accountCategory models.AccountCategory) uuid.UUID {
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

	rows, err := db.NamedQuery(query, accountCategory)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&accountCategory.AccountCategoryUUID)
		logError(err)
	}

	return accountCategory.AccountCategoryUUID
}

// CreateAccount creates an Account in db.
func (accountRepo *AccountRepository) CreateAccount(db *sqlx.DB, account models.Account) uuid.UUID {
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

	rows, err := db.NamedQuery(query, account)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&account.AccountUUID)
		logError(err)
	}

	return account.AccountUUID
}

// GetAccountCategory retrieves an AccountCategory from db.
func (accountRepo *AccountRepository) GetAccountCategory(db *sqlx.DB, accountCategoryUUID uuid.UUID) (accountCategory models.AccountCategory) {
	query := `
	SELECT
		account_category_uuid,
		name,
		description
	FROM account_category
	WHERE
		account_category_uuid = $1;`

	err := db.Get(&accountCategory, query, accountCategoryUUID.String())
	logError(err)

	return accountCategory
}

// GetAccountCategories retrieves AccountCategorys from db.
func (accountRepo *AccountRepository) GetAccountCategories(db *sqlx.DB) (accountCategories []models.AccountCategory) {
	query := `
	SELECT
		account_category_uuid,
		name,
		description
	FROM account_category;`

	err := db.Select(&accountCategories, query)
	logError(err)

	return accountCategories
}

// GetAccount retrieves an Account from db.
func (accountRepo *AccountRepository) GetAccount(db *sqlx.DB, accountUUID uuid.UUID) (account models.Account) {
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

	err := db.Get(&account, query, accountUUID.String())
	logError(err)

	return account
}

// GetAccounts retrieves Accounts from db.
func (accountRepo *AccountRepository) GetAccounts(db *sqlx.DB) (accounts []models.Account) {
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

	err := db.Select(&accounts, query)
	logError(err)

	return accounts
}

// UpdateAccountCategory updates an AccountCategory in db.
func (accountRepo *AccountRepository) UpdateAccountCategory(db *sqlx.DB, accountCategory models.AccountCategory) {
	query := `
	UPDATE account_category
	SET
		name = :name,
		description = :description
	WHERE
		account_category_uuid = :account_category_uuid;`

	_, err := db.NamedExec(query, accountCategory)
	logError(err)
}

// UpdateAccount updates an Account in db.
func (accountRepo *AccountRepository) UpdateAccount(db *sqlx.DB, account models.Account) {
	query := `
	UPDATE account
	SET
		account_category_uuid = :account_category.account_category_uuid,
		name = :name,
		description = :description,
		amount = :amount
	WHERE
		account_uuid = :account_uuid;`

	_, err := db.NamedExec(query, account)
	logError(err)
}

// DeleteAccountCategory deletes an AccountCategory from db.
func (accountRepo *AccountRepository) DeleteAccountCategory(db *sqlx.DB, accountCategoryUUID uuid.UUID) {
	query := `
	DELETE FROM account_category
	WHERE
		account_category_uuid = $1;`

	_, err := db.Exec(query, accountCategoryUUID.String())
	logError(err)
}

// DeleteAccount deletes an Account from db.
func (accountRepo *AccountRepository) DeleteAccount(db *sqlx.DB, accountUUID uuid.UUID) {
	query := `
	DELETE FROM account
	WHERE
		account_uuid = $1;`

	_, err := db.Exec(query, accountUUID.String())
	logError(err)
}
