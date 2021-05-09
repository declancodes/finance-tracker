package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AssetCategory is a type of asset a Fund has.
type AssetCategory struct {
	ID          uuid.UUID `json:"uuid,omitEmpty" db:"asset_category_uuid"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}

// Fund is a traded entity in which money can be invested.
type Fund struct {
	ID           uuid.UUID       `json:"uuid,omitEmpty" db:"fund_uuid"`
	Category     AssetCategory   `json:"category" db:"asset_category"`
	Name         string          `json:"name" db:"name"`
	TickerSymbol string          `json:"tickerSymbol" db:"ticker_symbol"`
	SharePrice   decimal.Decimal `json:"sharePrice" db:"share_price"`
	ExpenseRatio decimal.Decimal `json:"expenseRatio" db:"expense_ratio"`
	IsPrivate    bool            `json:"isPrivate" db:"is_private"`
}

// Holding is a position in a fund.
type Holding struct {
	ID               uuid.UUID       `json:"uuid,omitEmpty" db:"holding_uuid"`
	Account          Account         `json:"account" db:"account"`
	Fund             Fund            `json:"fund" db:"fund"`
	Shares           decimal.Decimal `json:"shares" db:"shares"`
	Value            decimal.Decimal `json:"value"`
	EffectiveExpense decimal.Decimal `json:"effectiveExpense"`
}

// PreviousPrice is the price for Date.
// It is always a day behind from the current date.
type PreviousPrice struct {
	Symbol         string          `json:"symbol"`
	Date           JSONTime        `json:"date"`
	Open           decimal.Decimal `json:"open"`
	Close          decimal.Decimal `json:"close"`
	High           decimal.Decimal `json:"high"`
	Low            decimal.Decimal `json:"low"`
	Volume         decimal.Decimal `json:"volume"`
	UOpen          decimal.Decimal `json:"uOpen"`
	UClose         decimal.Decimal `json:"uClose"`
	UHigh          decimal.Decimal `json:"uHigh"`
	ULow           decimal.Decimal `json:"uLow"`
	UVolume        decimal.Decimal `json:"uVolume"`
	Change         decimal.Decimal `json:"change"`
	ChangePercent  decimal.Decimal `json:"changePercent"`
	ChangeOverTime decimal.Decimal `json:"changeOverTime"`
}

// Marshaler is the interface used to marshal time.Time into the JSONTimeLayout custom date format.
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// Unmarshaler is the interface used to unmarshal time.Time in the JSONTimeLayout custom date format.
type Unmarshaler interface {
	UnmarshalJSON(bs []byte) error
}

// JSONTimeLayout is the custom time format for JSONTime.
const JSONTimeLayout = "2006-01-02"

// JSONTime is a custom time to format from JSON that is not the standard RFC.3339 format.
type JSONTime struct {
	time.Time
}

// MarshalJSON is used to marshal time.Time into the JSONTimeLayout custom date format.
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(JSONTimeLayout))
	return []byte(stamp), nil
}

// UnmarshalJSON is used to unmarshal time.Time in the JSONTimeLayout custom date format.
func (t JSONTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return nil
	}

	time, err := time.Parse(JSONTimeLayout, s)
	t.Time = time
	return err
}
