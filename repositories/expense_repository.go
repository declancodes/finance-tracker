package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseRepository is the means for interacting with Expense storage.
type ExpenseRepository struct{}

// CreateExpenseCategory creates an ExpenseCategory in db.
func (r *ExpenseRepository) CreateExpenseCategory(db *sqlx.DB, ec models.ExpenseCategory) uuid.UUID {
	query := `
	INSERT INTO expense_category (
		expense_category_uuid,
		name,
		description
	)
	VALUES (
		:expense_category_uuid,
		:name,
		:description
	)
	RETURNING expense_category_uuid;`

	rows, err := db.NamedQuery(query, ec)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&ec.ExpenseCategoryUUID)
		logError(err)
	}

	return ec.ExpenseCategoryUUID
}

// CreateExpense creates an Expense in db.
func (r *ExpenseRepository) CreateExpense(db *sqlx.DB, e models.Expense) uuid.UUID {
	query := `
	INSERT INTO expense (
		expense_uuid,
		expense_category_uuid,
		name,
		description,
		amount,
		date_incurred
	)
	VALUES (
		:expense_uuid,
		:expense_category.expense_category_uuid,
		:name,
		:description,
		:amount,
		:date_incurred
	)
	RETURNING expense_uuid;`

	rows, err := db.NamedQuery(query, e)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&e.ExpenseUUID)
		logError(err)
	}

	return e.ExpenseUUID
}

// GetExpenseCategory retrieves an ExpenseCategory from db.
func (r *ExpenseRepository) GetExpenseCategory(db *sqlx.DB, ecUUID uuid.UUID) (ec models.ExpenseCategory) {
	query := `
	SELECT
		expense_category_uuid,
		name,
		description
	FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	err := db.Get(&ec, query, ecUUID.String())
	logError(err)

	return ec
}

// GetExpenseCategories retrieves ExpenseCategorys from db.
func (r *ExpenseRepository) GetExpenseCategories(db *sqlx.DB) (ecs []models.ExpenseCategory) {
	query := `
	SELECT
		expense_category_uuid,
		name,
		description
	FROM expense_category;`

	err := db.Select(&ecs, query)
	logError(err)

	return ecs
}

// GetExpense retrieves an Expense from db.
func (r *ExpenseRepository) GetExpense(db *sqlx.DB, eUUID uuid.UUID) (e models.Expense) {
	query := `
	SELECT
		expense.expense_uuid,
		expense_category.expense_category_uuid AS "expense_category.expense_category_uuid",
		expense_category.name AS "expense_category.name",
		expense_category.description AS "expense_category.description",
		expense.name,
		expense.description,
		expense.amount,
		expense.date_incurred
	FROM expense
	INNER JOIN expense_category
		ON expense.expense_category_uuid = expense_category.expense_category_uuid
	WHERE
		expense.expense_uuid = $1;`

	err := db.Get(&e, query, eUUID.String())
	logError(err)

	return e
}

// GetExpenses retrieves Expenses from db.
func (r *ExpenseRepository) GetExpenses(db *sqlx.DB) (es []models.Expense) {
	query := `
	SELECT
		expense.expense_uuid,
		expense_category.expense_category_uuid AS "expense_category.expense_category_uuid",
		expense_category.name AS "expense_category.name",
		expense_category.description AS "expense_category.description",
		expense.name,
		expense.description,
		expense.amount,
		expense.date_incurred
	FROM expense
	INNER JOIN expense_category
		ON expense.expense_category_uuid = expense_category.expense_category_uuid;`

	err := db.Select(&es, query)
	logError(err)

	return es
}

// UpdateExpenseCategory updates an ExpenseCategory in db.
func (r *ExpenseRepository) UpdateExpenseCategory(db *sqlx.DB, ec models.ExpenseCategory) {
	query := `
	UPDATE expense_category
	SET
		name = :name,
		description = :description
	WHERE
		expense_category_uuid = :expense_category_uuid;`

	_, err := db.NamedExec(query, ec)
	logError(err)
}

// UpdateExpense updates an Expense in db.
func (r *ExpenseRepository) UpdateExpense(db *sqlx.DB, e models.Expense) {
	query := `
	UPDATE expense
	SET
		expense_category_uuid = :expense_category.expense_category_uuid,
		name = :name,
		description = :description,
		amount = :amount,
		date_incurred = :date_incurred
	WHERE
		expense_uuid = :expense_uuid;`

	_, err := db.NamedExec(query, e)
	logError(err)
}

// DeleteExpenseCategory deletes an ExpenseCategory from db.
func (r *ExpenseRepository) DeleteExpenseCategory(db *sqlx.DB, ecUUID uuid.UUID) {
	query := `
	DELETE FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	_, err := db.Exec(query, ecUUID.String())
	logError(err)
}

// DeleteExpense deletes an Expense from db.
func (r *ExpenseRepository) DeleteExpense(db *sqlx.DB, eUUID uuid.UUID) {
	query := `
	DELETE FROM expense
	WHERE
		expense_uuid = $1;`

	_, err := db.Exec(query, eUUID.String())
	logError(err)
}
