package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AccountCategory is a type of Account.
// This might be something like 'Checking', 'Savings', 'IRA', etc.
type AccountCategory struct {
	ID          uuid.UUID `json:"uuid,omitEmpty" db:"account_category_uuid"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

// Account is a vehicle in which money is stored.
type Account struct {
	ID          uuid.UUID       `json:"uuid,omitEmpty" db:"account_uuid"`
	Category    AccountCategory `json:"category" db:"account_category"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
}

// Contribution is a payment made from Income to an Account.
// This might be something like a savings Account deposit, retirement plan deposit, etc.
type Contribution struct {
	ID          uuid.UUID       `json:"uuid,omitEmpty" db:"contribution_uuid"`
	Account     Account         `json:"account" db:"account"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Date        time.Time       `json:"date" db:"date_made"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
}

// Income is a payment taken from work or investment returns.
// This might be something like a paycheck deposit, a savings Account interest payment, income from investment, etc.
type Income struct {
	ID          uuid.UUID       `json:"uuid,omitEmpty" db:"income_uuid"`
	Account     Account         `json:"account" db:"account"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Date        time.Time       `json:"date" db:"date_made"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
}
