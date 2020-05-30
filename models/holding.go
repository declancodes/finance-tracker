package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Holding is a position in a fund.
type Holding struct {
	ID      uuid.UUID       `json:"uuid,omitEmpty" db:"holding_uuid"`
	Account Account         `json:"account" db:"account"`
	Fund    Fund            `json:"fund" db:"fund"`
	Shares  decimal.Decimal `json:"shares" db:"shares"`
	Value   decimal.Decimal `json:"value"`
}
