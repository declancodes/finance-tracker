package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseRepository is the means for interacting with Expense storage.
type ExpenseRepository struct{}

// CreateExpenseCategory creates an ExpenseCategory in db.
func (expenseRepo *ExpenseRepository) CreateExpenseCategory(db *sqlx.DB, expenseCategory models.ExpenseCategory) uuid.UUID {
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

	rows, err := db.NamedQuery(query, expenseCategory)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&expenseCategory.ExpenseCategoryUUID)
		logError(err)
	}

	return expenseCategory.ExpenseCategoryUUID
}

// CreateExpense creates an Expense in db.
func (expenseRepo *ExpenseRepository) CreateExpense(db *sqlx.DB, expense models.Expense) uuid.UUID {
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

	rows, err := db.NamedQuery(query, expense)
	logError(err)

	for rows.Next() {
		err = rows.Scan(&expense.ExpenseUUID)
		logError(err)
	}

	return expense.ExpenseUUID
}

// GetExpenseCategory retrieves an ExpenseCategory from db.
func (expenseRepo *ExpenseRepository) GetExpenseCategory(db *sqlx.DB, expenseCategoryUUID uuid.UUID) (expenseCategory models.ExpenseCategory) {
	query := `
	SELECT
		expense_category_uuid,
		name,
		description
	FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	err := db.Get(&expenseCategory, query, expenseCategoryUUID.String())
	logError(err)

	return expenseCategory
}

// GetExpenseCategories retrieves ExpenseCategorys from db.
func (expenseRepo *ExpenseRepository) GetExpenseCategories(db *sqlx.DB) (expenseCategories []models.ExpenseCategory) {
	query := `
	SELECT
		expense_category_uuid,
		name,
		description
	FROM expense_category;`

	err := db.Select(&expenseCategories, query)
	logError(err)

	return expenseCategories
}

// GetExpense retrieves an Expense from db.
func (expenseRepo *ExpenseRepository) GetExpense(db *sqlx.DB, expenseUUID uuid.UUID) (expense models.Expense) {
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

	err := db.Get(&expense, query, expenseUUID.String())
	logError(err)

	return expense
}

// GetExpenses retrieves Expenses from db.
func (expenseRepo *ExpenseRepository) GetExpenses(db *sqlx.DB) (expenses []models.Expense) {
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

	err := db.Select(&expenses, query)
	logError(err)

	return expenses
}

// UpdateExpenseCategory updates an ExpenseCategory in db.
func (expenseRepo *ExpenseRepository) UpdateExpenseCategory(db *sqlx.DB, expenseCategory models.ExpenseCategory) {
	query := `
	UPDATE expense_category
	SET
		name = :name,
		description = :description
	WHERE
		expense_category_uuid = :expense_category_uuid;`

	_, err := db.NamedExec(query, expenseCategory)
	logError(err)
}

// UpdateExpense updates an Expense in db.
func (expenseRepo *ExpenseRepository) UpdateExpense(db *sqlx.DB, expense models.Expense) {
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

	_, err := db.NamedExec(query, expense)
	logError(err)
}

// DeleteExpenseCategory deletes an ExpenseCategory from db.
func (expenseRepo *ExpenseRepository) DeleteExpenseCategory(db *sqlx.DB, expenseCategoryUUID uuid.UUID) {
	query := `
	DELETE FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	_, err := db.Exec(query, expenseCategoryUUID.String())
	logError(err)
}

// DeleteExpense deletes an Expense from db.
func (expenseRepo *ExpenseRepository) DeleteExpense(db *sqlx.DB, expenseUUID uuid.UUID) {
	query := `
	DELETE FROM expense
	WHERE
		expense_uuid = $1;`

	_, err := db.Exec(query, expenseUUID.String())
	logError(err)
}
