package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExpenseCategory is .
type ExpenseCategory struct {
	ExpenseCategoryUUID uuid.UUID `json:"expenseCategoryUuid,omitEmpty" db:"expense_category_uuid"`
	Name                string    `json:"name" db:"name"`
	Description         string    `json:"description" db:"description"`
}

// Expense is .
type Expense struct {
	ExpenseUUID     uuid.UUID       `json:"expenseUuid,omitEmpty" db:"expense_uuid"`
	ExpenseCategory ExpenseCategory `json:"expenseCategory" db:"expense_category"`
	Name            string          `json:"name" db:"name"`
	Description     string          `json:"description" db:"description"`
	Date            time.Time       `json:"date" db:"date_incurred"`
	Amount          decimal.Decimal `json:"amount" db:"amount"`
}
