package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseController is the means for interacting with Expense entities from an http router.
type ExpenseController struct{}

var expenseRepo = repositories.ExpenseRepository{}

func badRequestExpenseCategory(w http.ResponseWriter, err error) {
	badRequestModel(w, "expense category", err)
}

func badRequestExpense(w http.ResponseWriter, err error) {
	badRequestModel(w, "expense", err)
}

func errorExecutingExpenseCategory(w http.ResponseWriter, err error) {
	errorExecuting(w, "expense category", err)
}

func errorExecutingExpense(w http.ResponseWriter, err error) {
	errorExecuting(w, "expense", err)
}

// CreateExpenseCategory creates an ExpenseCategory based on the r *http.Request Body.
func (c *ExpenseController) CreateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ec *models.ExpenseCategory
		err := json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			badRequestExpenseCategory(w, err)
			return
		}

		ec.ID = uuid.New()
		ecIDs, err := expenseRepo.CreateExpenseCategories(db, []*models.ExpenseCategory{ec})
		if err != nil {
			errorCreating(w, "expense category", err)
			return
		}

		created(w, ecIDs[0])
	}
}

// CreateExpense creates an Expense based on the r *http.Request Body.
func (c *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e *models.Expense
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			badRequestExpense(w, err)
			return
		}

		e.ID = uuid.New()
		eIDs, err := expenseRepo.CreateExpenses(db, []*models.Expense{e})
		if err != nil {
			errorCreating(w, "expense", err)
			return
		}

		created(w, eIDs[0])
	}
}

// GetExpenseCategory gets an ExpenseCategory.
func (c *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		ec, err := expenseRepo.GetExpenseCategory(db, ecID)
		if err != nil {
			errorExecutingExpenseCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(ec)
		logError(err)
	}
}

// GetExpenseCategories gets ExpenseCategory entities.
func (c *ExpenseController) GetExpenseCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecs, err := expenseRepo.GetExpenseCategories(db, getFilters(r))
		if err != nil {
			errorExecutingExpenseCategory(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(ecs)
		logError(err)
	}
}

// GetExpense gets an Expense.
func (c *ExpenseController) GetExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		e, err := expenseRepo.GetExpense(db, eID)
		if err != nil {
			errorExecutingExpense(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(e)
		logError(err)
	}
}

// GetExpenses gets Expense entities.
func (c *ExpenseController) GetExpenses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		es, err := expenseRepo.GetExpenses(db, getFilters(r))

		if err != nil {
			errorExecutingExpense(w, err)
			return
		}

		addJSONContentHeader(w)
		err = json.NewEncoder(w).Encode(es)
		logError(err)
	}
}

// UpdateExpenseCategory updates an ExpenseCategory.
func (c *ExpenseController) UpdateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var ec *models.ExpenseCategory
		err = json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			badRequestExpenseCategory(w, err)
			return
		}

		ec.ID = ecID
		err = expenseRepo.UpdateExpenseCategory(db, ec)
		if err != nil {
			errorExecutingExpenseCategory(w, err)
			return
		}

		updated(w, ec.ID)
	}
}

// UpdateExpense updates an Expense.
func (c *ExpenseController) UpdateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eID, err := getID(r)
		if err != nil {
			badRequestID(w, err)
			return
		}

		var e *models.Expense
		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			badRequestExpense(w, err)
			return
		}

		e.ID = eID
		err = expenseRepo.UpdateExpense(db, e)
		if err != nil {
			errorExecutingExpense(w, err)
			return
		}

		updated(w, e.ID)
	}
}

// DeleteExpenseCategory deletes an ExpenseCategory.
func (c *ExpenseController) DeleteExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "expense category", expenseRepo.DeleteExpenseCategory)
	}
}

// DeleteExpense deletes an Expense.
func (c *ExpenseController) DeleteExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delete(w, r, db, "expense", expenseRepo.DeleteExpense)
	}
}
