package repositories

import (
	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseRepository is the means for interacting with Expense storage.
type ExpenseRepository struct{}

// CreateExpenseCategories creates ExpenseCategory entities in db.
func (r *ExpenseRepository) CreateExpenseCategories(db *sqlx.DB, ecs []*models.ExpenseCategory) ([]uuid.UUID, error) {
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

	IDs, err := createAndGetIDs(db, query, ecs)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// CreateExpenses creates Expense entities in db.
func (r *ExpenseRepository) CreateExpenses(db *sqlx.DB, es []*models.Expense) ([]uuid.UUID, error) {
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

	IDs, err := createAndGetIDs(db, query, es)
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

// GetExpenseCategory retrieves ExpenseCategory with ecID from db.
func (r *ExpenseRepository) GetExpenseCategory(db *sqlx.DB, ecID uuid.UUID) (*models.ExpenseCategory, error) {
	mValues := map[string]interface{}{
		"expense_category": ecID.String(),
	}

	ecs, err := r.GetExpenseCategories(db, mValues)
	if err != nil {
		return nil, err
	}

	return ecs[0], nil
}

// GetExpenseCategories retrieves ExpenseCategory entities from db.
// Filters for ExpenseCategory retrieval are applied to the query based on the key-value pairs in mValues.
func (r *ExpenseRepository) GetExpenseCategories(db *sqlx.DB, mValues map[string]interface{}) ([]*models.ExpenseCategory, error) {
	query := `
	SELECT
		expense_category.expense_category_uuid,
		expense_category.name,
		expense_category.description
	FROM expense_category`

	mFilters := map[string]string{
		"expense_category": "expense_category.expense_category_uuid = ",
	}

	q, args, err := getGetQueryAndValues(query, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var ecs []*models.ExpenseCategory
	err = db.Select(&ecs, q, args...)
	if err != nil {
		return nil, err
	}
	return ecs, nil
}

// GetExpense retrieves Expense with eID from db.
func (r *ExpenseRepository) GetExpense(db *sqlx.DB, eID uuid.UUID) (*models.Expense, error) {
	mValues := map[string]interface{}{
		"expense": eID.String(),
	}

	es, err := r.GetExpenses(db, mValues)
	if err != nil {
		return nil, err
	}
	return es[0], nil
}

// GetExpenses retrieves Expense entities from db.
// Filters for Expense retrieval are applied to the query based on the key-value pairs in mValues.
func (r *ExpenseRepository) GetExpenses(db *sqlx.DB, mValues map[string]interface{}) ([]*models.Expense, error) {
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
		ON expense.expense_category_uuid = expense_category.expense_category_uuid`

	mFilters := map[string]string{
		"expense":    "expense.expense_uuid = ",
		"categories": "expense_category.name IN ",
		"start":      "expense.date_incurred >= ",
		"end":        "expense.date_incurred <= ",
	}

	q, args, err := getGetQueryAndValues(query, mValues, mFilters)
	if err != nil {
		return nil, err
	}

	var es []*models.Expense
	err = db.Select(&es, q, args...)
	if err != nil {
		return nil, err
	}
	return es, nil
}

// UpdateExpenseCategory updates an ExpenseCategory in db.
func (r *ExpenseRepository) UpdateExpenseCategory(db *sqlx.DB, ec *models.ExpenseCategory) error {
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
func (r *ExpenseRepository) UpdateExpense(db *sqlx.DB, e *models.Expense) error {
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
func (r *ExpenseRepository) DeleteExpenseCategory(db *sqlx.DB, ecID uuid.UUID) error {
	query := `
	DELETE FROM expense_category
	WHERE
		expense_category_uuid = $1;`

	return deleteEntity(db, query, ecID)
}

// DeleteExpense deletes an Expense from db.
func (r *ExpenseRepository) DeleteExpense(db *sqlx.DB, eID uuid.UUID) error {
	query := `
	DELETE FROM expense
	WHERE
		expense_uuid = $1;`

	return deleteEntity(db, query, eID)
}
