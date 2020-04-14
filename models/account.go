package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AccountCategory is .
type AccountCategory struct {
	AccountCategoryUUID uuid.UUID `json:"accountCategoryUuid,omitEmpty" db:"account_category_uuid"`
	Name                string    `json:"name" db:"name"`
	Description         string    `json:"description" db:"description"`
}

// Account is .
type Account struct {
	AccountUUID     uuid.UUID       `json:"accountUuid,omitEmpty" db:"account_uuid"`
	AccountCategory AccountCategory `json:"accountCategory" db:"account_category"`
	Name            string          `json:"name" db:"name"`
	Description     string          `json:"description" db:"description"`
	Amount          decimal.Decimal `json:"amount" db:"amount"`
}
