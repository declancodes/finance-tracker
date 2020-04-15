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

// CreateExpenseCategory .
func (expenseController *ExpenseController) CreateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory
		expenseCategoryUUID, _ := uuid.NewUUID()

		err := json.NewDecoder(r.Body).Decode(&expenseCategory)
		logError(err)

		expenseCategory.ExpenseCategoryUUID = expenseCategoryUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseCategoryUUID = expenseRepo.CreateExpenseCategory(db, expenseCategory)

		err = json.NewEncoder(w).Encode(expenseCategoryUUID)
		logError(err)
	}
}

// CreateExpense .
func (expenseController *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID, _ := uuid.NewUUID()

		err := json.NewDecoder(r.Body).Decode(&expense)
		logError(err)

		expense.ExpenseUUID = expenseUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseUUID = expenseRepo.CreateExpense(db, expense)

		err = json.NewEncoder(w).Encode(expenseUUID)
		logError(err)
	}
}

// GetExpenseCategory .
func (expenseController *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory
		expenseCategoryUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expenseCategory = expenseRepo.GetExpenseCategory(db, expenseCategoryUUID)

		err := json.NewEncoder(w).Encode(expenseCategory)
		logError(err)
	}
}

// GetExpenseCategories .
func (expenseController *ExpenseController) GetExpenseCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseRepo := repositories.ExpenseRepository{}
		expenseCategories := expenseRepo.GetExpenseCategories(db)

		err := json.NewEncoder(w).Encode(expenseCategories)
		logError(err)
	}
}

// GetExpense .
func (expenseController *ExpenseController) GetExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expense = expenseRepo.GetExpense(db, expenseUUID)

		err := json.NewEncoder(w).Encode(expense)
		logError(err)
	}
}

// GetExpenses .
func (expenseController *ExpenseController) GetExpenses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseRepo := repositories.ExpenseRepository{}
		expenses := expenseRepo.GetExpenses(db)

		err := json.NewEncoder(w).Encode(expenses)
		logError(err)
	}
}

// UpdateExpenseCategory .
func (expenseController *ExpenseController) UpdateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory
		expenseCategoryUUID := getUUID(r)

		err := json.NewDecoder(r.Body).Decode(&expenseCategory)
		logError(err)

		expenseCategory.ExpenseCategoryUUID = expenseCategoryUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.UpdateExpenseCategory(db, expenseCategory)

		err = json.NewEncoder(w).Encode(expenseCategory)
		logError(err)
	}
}

// UpdateExpense .
func (expenseController *ExpenseController) UpdateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID := getUUID(r)

		err := json.NewDecoder(r.Body).Decode(&expense)
		logError(err)

		expense.ExpenseUUID = expenseUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.UpdateExpense(db, expense)

		err = json.NewEncoder(w).Encode(expense)
		logError(err)
	}
}

// DeleteExpenseCategory .
func (expenseController *ExpenseController) DeleteExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseCategoryUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.DeleteExpenseCategory(db, expenseCategoryUUID)
	}
}

// DeleteExpense .
func (expenseController *ExpenseController) DeleteExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.DeleteExpense(db, expenseUUID)
	}
}
