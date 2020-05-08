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

	rows, err := db.NamedQuery(query, ec)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ec.ID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return ec.ID, nil
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

	rows, err := db.NamedQuery(query, e)
	if err != nil {
		return uuid.Nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&e.ID)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return e.ID, nil
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
	query := fmt.Sprintf(`
	%s
	WHERE
		expense.expense_uuid = $1;`, getExpensesQuery)

	err = db.Get(&e, query, eUUID.String())
	return e, err
}

// GetExpenses retrieves Expense entities from db.
// Filters for Expense retrieval are applied to the query based on the key-value pairs in mValues.
func (r *ExpenseRepository) GetExpenses(db *sqlx.DB, mValues map[string]interface{}) (es []models.Expense, err error) {
	mFilters := map[string]string{
		"category": "expense_category.name = ",
		"start":    "expense.date_incurred >= ",
		"end":      "expense.date_incurred <= ",
	}

	clauses, values, err := buildQueryClauses(mValues, mFilters)
	if err != nil {
		return es, err
	}

	query := fmt.Sprintf("%s %s", getExpensesQuery, clauses)

	err = db.Select(&es, query, values...)
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

	res, err := db.NamedExec(query, ec)

	_, err = getExecuted(res, err)
	return err
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

	res, err := db.NamedExec(query, e)

	_, err = getExecuted(res, err)
	return err
}

// DeleteExpenseCategory deletes an ExpenseCategory from db.
func (r *ExpenseRepository) DeleteExpenseCategory(db *sqlx.DB, ecUUID uuid.UUID) error {
	query := `
	DELETE FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	res, err := db.Exec(query, ecUUID.String())

	_, err = getExecuted(res, err)
	return err
}

// DeleteExpense deletes an Expense from db.
func (r *ExpenseRepository) DeleteExpense(db *sqlx.DB, eUUID uuid.UUID) error {
	query := `
	DELETE FROM expense
	WHERE
		expense_uuid = $1;`

	res, err := db.Exec(query, eUUID.String())

	_, err = getExecuted(res, err)
	return err
}
