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

		json.NewDecoder(r.Body).Decode(&expenseCategory)
		expenseCategory.ExpenseCategoryUUID = expenseCategoryUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseCategoryUUID = expenseRepo.CreateExpenseCategory(db, expenseCategory)

		json.NewEncoder(w).Encode(expenseCategoryUUID)
	}
}

// CreateExpense .
func (expenseController *ExpenseController) CreateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID, _ := uuid.NewUUID()

		json.NewDecoder(r.Body).Decode(&expense)
		expense.ExpenseUUID = expenseUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseUUID = expenseRepo.CreateExpense(db, expense)

		json.NewEncoder(w).Encode(expenseUUID)
	}
}

// GetExpenseCategory .
func (expenseController *ExpenseController) GetExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory
		expenseCategoryUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expenseCategory = expenseRepo.GetExpenseCategory(db, expenseCategoryUUID)

		json.NewEncoder(w).Encode(expenseCategory)
	}
}

// GetExpenseCategories .
func (expenseController *ExpenseController) GetExpenseCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseRepo := repositories.ExpenseRepository{}
		expenseCategories := expenseRepo.GetExpenseCategories(db)

		json.NewEncoder(w).Encode(expenseCategories)
	}
}

// GetExpense .
func (expenseController *ExpenseController) GetExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID := getUUID(r)

		expenseRepo := repositories.ExpenseRepository{}
		expense = expenseRepo.GetExpense(db, expenseUUID)

		json.NewEncoder(w).Encode(expense)
	}
}

// GetExpenses .
func (expenseController *ExpenseController) GetExpenses(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expenseRepo := repositories.ExpenseRepository{}
		expenses := expenseRepo.GetExpenses(db)

		json.NewEncoder(w).Encode(expenses)
	}
}

// UpdateExpenseCategory .
func (expenseController *ExpenseController) UpdateExpenseCategory(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expenseCategory models.ExpenseCategory
		expenseCategoryUUID := getUUID(r)

		json.NewDecoder(r.Body).Decode(&expenseCategory)
		expenseCategory.ExpenseCategoryUUID = expenseCategoryUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.UpdateExpenseCategory(db, expenseCategory)

		json.NewEncoder(w).Encode(expenseCategory)
	}
}

// UpdateExpense .
func (expenseController *ExpenseController) UpdateExpense(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var expense models.Expense
		expenseUUID := getUUID(r)

		json.NewDecoder(r.Body).Decode(&expense)
		expense.ExpenseUUID = expenseUUID

		expenseRepo := repositories.ExpenseRepository{}
		expenseRepo.UpdateExpense(db, expense)

		json.NewEncoder(w).Encode(expense)
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
