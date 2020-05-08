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
		var ec models.ExpenseCategory

		err := json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			badRequestExpenseCategory(w, err)
			return
		}

		ec.ID, _ = uuid.NewUUID()
		ecUUID, err := expenseRepo.CreateExpenseCategory(db, ec)
		if err != nil {
			errorCreating(w, "expense category", err)
			return
		}

		created(w, ecUUID)
	}
}

// CreateExpense creates an Expense based on the r *http.Request Body.
func (c *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Expense
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			badRequestExpense(w, err)
			return
		}

		e.ID, _ = uuid.NewUUID()
		eUUID, err := expenseRepo.CreateExpense(db, e)
		if err != nil {
			errorCreating(w, "expense", err)
			return
		}

		created(w, eUUID)
	}
}

// GetExpenseCategory gets an ExpenseCategory.
func (c *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		ec, err := expenseRepo.GetExpenseCategory(db, ecUUID)
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
		ecs, err := expenseRepo.GetExpenseCategories(db)
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
		eUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		e, err := expenseRepo.GetExpense(db, eUUID)
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
		q := r.URL.Query()
		catName := q.Get("category")
		start := getTime(q.Get("start"))
		end := getTime(q.Get("end"))

		mValues := make(map[string]interface{})
		if catName != "" {
			mValues["category"] = catName
		}
		if !start.IsZero() {
			mValues["start"] = start
		}
		if !end.IsZero() {
			mValues["end"] = end
		}

		es, err := expenseRepo.GetExpenses(db, mValues)

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
		ecUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var ec models.ExpenseCategory
		err = json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			badRequestExpenseCategory(w, err)
			return
		}

		ec.ID = ecUUID
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
		eUUID, err := getUUID(r)
		if err != nil {
			badRequestUUID(w, err)
			return
		}

		var e models.Expense
		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			badRequestExpense(w, err)
			return
		}

		e.ID = eUUID
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
