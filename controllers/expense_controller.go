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
func (expenseController *ExpenseController) CreateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory

		err := json.NewDecoder(r.Body).Decode(&expenseCategory)
		if err != nil {
			writeHeaderForBadRequest(w, "invalid expense category", err)
			return
		}

		expenseCategory.ExpenseCategoryUUID, _ = uuid.NewUUID()
		expenseCategoryUUID := expenseRepo.CreateExpenseCategory(db, expenseCategory)

		err = json.NewEncoder(w).Encode(expenseCategoryUUID)
		logError(err)
	}
}

// CreateExpense .
func (expenseController *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense

		err := json.NewDecoder(r.Body).Decode(&expense)
		if err != nil {
			writeHeaderForBadRequest(w, "invalid expense", err)
			return
		}

		expense.ExpenseUUID, _ = uuid.NewUUID()
		expenseUUID := expenseRepo.CreateExpense(db, expense)

		err = json.NewEncoder(w).Encode(expenseUUID)
		logError(err)
	}
}

// GetExpenseCategory .
func (expenseController *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseCategoryUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expenseCategory := expenseRepo.GetExpenseCategory(db, expenseCategoryUUID)

		err = json.NewEncoder(w).Encode(expenseCategory)
		logError(err)
	}
}

// GetExpenseCategories .
func (expenseController *ExpenseController) GetExpenseCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseCategories := expenseRepo.GetExpenseCategories(db)

		err := json.NewEncoder(w).Encode(expenseCategories)
		logError(err)
	}
}

// GetExpense .
func (expenseController *ExpenseController) GetExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expense := expenseRepo.GetExpense(db, expenseUUID)

		err = json.NewEncoder(w).Encode(expense)
		logError(err)
	}
}

// GetExpenses .
func (expenseController *ExpenseController) GetExpenses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenses := expenseRepo.GetExpenses(db)

		err := json.NewEncoder(w).Encode(expenses)
		logError(err)
	}
}

// UpdateExpenseCategory .
func (expenseController *ExpenseController) UpdateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseCategoryUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var expenseCategory models.ExpenseCategory

		err = json.NewDecoder(r.Body).Decode(&expenseCategory)
		if err != nil {
			writeHeaderForBadRequest(w, "invalid expense category", err)
			return
		}

		expenseCategory.ExpenseCategoryUUID = expenseCategoryUUID
		expenseRepo.UpdateExpenseCategory(db, expenseCategory)

		err = json.NewEncoder(w).Encode(expenseCategory)
		logError(err)
	}
}

// UpdateExpense .
func (expenseController *ExpenseController) UpdateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		var expense models.Expense

		err = json.NewDecoder(r.Body).Decode(&expense)
		if err != nil {
			writeHeaderForBadRequest(w, "invalid expense", err)
			return
		}

		expense.ExpenseUUID = expenseUUID
		expenseRepo.UpdateExpense(db, expense)

		err = json.NewEncoder(w).Encode(expense)
		logError(err)
	}
}

// DeleteExpenseCategory .
func (expenseController *ExpenseController) DeleteExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseCategoryUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expenseRepo.DeleteExpenseCategory(db, expenseCategoryUUID)
	}
}

// DeleteExpense .
func (expenseController *ExpenseController) DeleteExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseUUID, err := getUUID(r)
		if err != nil {
			writeHeaderForBadRequestUUID(w, err)
			return
		}

		expenseRepo.DeleteExpense(db, expenseUUID)
	}
}
