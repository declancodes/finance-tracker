package repositories

import (
	"fmt"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseRepository is the means for interacting with Expense storage.
type ExpenseRepository struct{}

const (
	getExpenseCategoriesQuery = `
	SELECT
		expense_category_uuid,
		name,
		description
	FROM expense_category`

	getExpensesQuery = `
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
		ON expense.expense_category_uuid = expense_category.expense_category_uuid`
)

// CreateExpenseCategory creates an ExpenseCategory in db.
func (r *ExpenseRepository) CreateExpenseCategory(db *sqlx.DB, ec models.ExpenseCategory) (uuid.UUID, error) {
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

	return createAndGetUUID(db, query, ec)
}

// CreateExpense creates an Expense in db.
func (r *ExpenseRepository) CreateExpense(db *sqlx.DB, e models.Expense) (uuid.UUID, error) {
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

	return createAndGetUUID(db, query, e)
}

// GetExpenseCategory retrieves ExpenseCategory with ecUUID from db.
func (r *ExpenseRepository) GetExpenseCategory(db *sqlx.DB, ecUUID uuid.UUID) (ec models.ExpenseCategory, err error) {
	query := fmt.Sprintf(`
	%s
	WHERE
		expense_category_uuid = $1;`, getExpenseCategoriesQuery)

	err = db.Get(&ec, query, ecUUID.String())
	return ec, err
}

// GetExpenseCategories retrieves ExpenseCategory entities from db.
func (r *ExpenseRepository) GetExpenseCategories(db *sqlx.DB) (ecs []models.ExpenseCategory, err error) {
	query := fmt.Sprintf(`%s;`, getExpenseCategoriesQuery)

	err = db.Select(&ecs, query)
	return ecs, err
}

// GetExpense retrieves Expense with eUUID from db.
func (r *ExpenseRepository) GetExpense(db *sqlx.DB, eUUID uuid.UUID) (e models.Expense, err error) {
	mValues := map[string]interface{}{
		"expense": eUUID.String(),
	}

	es, err := r.GetExpenses(db, mValues)
	if err != nil {
		return e, err
	}
	if len(es) > 1 {
		return e, fmt.Errorf("more than one Expense with ID: %v", eUUID)
	}

	return es[0], nil
}

// GetExpenses retrieves Expense entities from db.
// Filters for Expense retrieval are applied to the query based on the key-value pairs in mValues.
func (r *ExpenseRepository) GetExpenses(db *sqlx.DB, mValues map[string]interface{}) (es []models.Expense, err error) {
	mFilters := map[string]string{
		"expense":    "expense.expense_uuid = ",
		"categories": "expense_category.name IN ",
		"start":      "expense.date_incurred >= ",
		"end":        "expense.date_incurred <= ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s %s", getExpensesQuery, clauses)

	q, args, err := sqlx.In(query, values...)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)

	err = db.Select(&es, q, args...)
	return es, err
}

// UpdateExpenseCategory updates an ExpenseCategory in db.
func (r *ExpenseRepository) UpdateExpenseCategory(db *sqlx.DB, ec models.ExpenseCategory) error {
	query := `
	UPDATE expense_category
	SET
		name = :name,
		description = :description
	WHERE
		expense_category_uuid = :expense_category_uuid;`

	return updateEntity(db, query, ec)
}

// UpdateExpense updates an Expense in db.
func (r *ExpenseRepository) UpdateExpense(db *sqlx.DB, e models.Expense) error {
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

	return updateEntity(db, query, e)
}

// DeleteExpenseCategory deletes an ExpenseCategory from db.
func (r *ExpenseRepository) DeleteExpenseCategory(db *sqlx.DB, ecUUID uuid.UUID) error {
	query := `
	DELETE FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	return deleteEntity(db, query, ecUUID)
}

// DeleteExpense deletes an Expense from db.
func (r *ExpenseRepository) DeleteExpense(db *sqlx.DB, eUUID uuid.UUID) error {
	query := `
	DELETE FROM expense
	WHERE
		expense_uuid = $1;`

	return deleteEntity(db, query, eUUID)
}
