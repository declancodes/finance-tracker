package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Fund is a traded entity in which money can be invested.
type Fund struct {
	ID           uuid.UUID       `json:"uuid,omitEmpty" db:"fund_uuid"`
	Name         string          `json:"name" db:"name"`
	TickerSymbol string          `json:"tickerSymbol" db:"ticker_symbol"`
	SharePrice   decimal.Decimal `json:"sharePrice" db:"share_price"`
}
