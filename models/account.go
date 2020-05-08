package models

import (
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
