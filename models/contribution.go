package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Contribution is a payment made to an Account.
// This might be something like a paycheck deposit, a savings Account deposit, retirement plan deposit, etc.
type Contribution struct {
	ID          uuid.UUID       `json:"uuid,omitEmpty" db:"contribution_uuid"`
	Account     Account         `json:"account" db:"account"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Date        time.Time       `json:"date" db:"date_made"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
}
