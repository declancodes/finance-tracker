package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExpenseCategory is a type of Expense.
// This might be something like 'Living', 'Travel', 'Groceries', etc.
type ExpenseCategory struct {
	ID          uuid.UUID `json:"uuid,omitEmpty" db:"expense_category_uuid"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

// Expense is an amount of money paid elsewhere.
// This might be something like a rent/mortgage payment, a grocery bill, entertainment purchase, etc.
type Expense struct {
	ID          uuid.UUID       `json:"uuid,omitEmpty" db:"expense_uuid"`
	Category    ExpenseCategory `json:"category" db:"expense_category"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Date        time.Time       `json:"date" db:"date_incurred"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
}
