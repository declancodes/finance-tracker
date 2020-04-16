package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DeclanCodes/finance-tracker/models"
	"github.com/DeclanCodes/finance-tracker/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExpenseController .
type ExpenseController struct{}

var expenseRepo = repositories.ExpenseRepository{}

// CreateExpenseCategory .
func (c *ExpenseController) CreateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ec models.ExpenseCategory

		err := json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			writeHeaderForBadRequestModel(w, "expense category", err)
			return
		}

		ec.ExpenseCategoryUUID, _ = uuid.NewUUID()
		ecUUID := expenseRepo.CreateExpenseCategory(db, ec)

		err = json.NewEncoder(w).Encode(ecUUID)
		logError(err)
	}
}

// CreateExpense .
func (c *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Expense

		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			writeHeaderForBadRequestModel(w, "expense", err)
			return
		}

		e.ExpenseUUID, _ = uuid.NewUUID()
		eUUID := expenseRepo.CreateExpense(db, e)

		err = json.NewEncoder(w).Encode(eUUID)
		logError(err)
	}
}

// GetExpenseCategory .
func (c *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		ec := expenseRepo.GetExpenseCategory(db, ecUUID)

		err = json.NewEncoder(w).Encode(ec)
		logError(err)
	}
}

// GetExpenseCategories .
func (c *ExpenseController) GetExpenseCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecs := expenseRepo.GetExpenseCategories(db)

		err := json.NewEncoder(w).Encode(ecs)
		logError(err)
	}
}

// GetExpense .
func (c *ExpenseController) GetExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		e := expenseRepo.GetExpense(db, eUUID)

		err = json.NewEncoder(w).Encode(e)
		logError(err)
	}
}

// GetExpenses .
func (c *ExpenseController) GetExpenses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		es := expenseRepo.GetExpenses(db)

		err := json.NewEncoder(w).Encode(es)
		logError(err)
	}
}

// UpdateExpenseCategory .
func (c *ExpenseController) UpdateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var ec models.ExpenseCategory

		err = json.NewDecoder(r.Body).Decode(&ec)
		if err != nil {
			writeHeaderForBadRequestModel(w, "expense category", err)
			return
		}

		ec.ExpenseCategoryUUID = ecUUID
		expenseRepo.UpdateExpenseCategory(db, ec)

		err = json.NewEncoder(w).Encode(ec)
		logError(err)
	}
}

// UpdateExpense .
func (c *ExpenseController) UpdateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var e models.Expense

		err = json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			writeHeaderForBadRequestModel(w, "expense", err)
			return
		}

		e.ExpenseUUID = eUUID
		expenseRepo.UpdateExpense(db, e)

		err = json.NewEncoder(w).Encode(e)
		logError(err)
	}
}

// DeleteExpenseCategory .
func (c *ExpenseController) DeleteExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ecUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expenseRepo.DeleteExpenseCategory(db, ecUUID)
	}
}

// DeleteExpense .
func (c *ExpenseController) DeleteExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expenseRepo.DeleteExpense(db, eUUID)
	}
}
