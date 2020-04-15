# Planning for finance-tracker

Expenses
  - Expenses and contributions to things.
  - Worry about savings ratio, not necessarily amounts in accounts

Investments
  - Worry about actual amounts in accounts and expense ratios.
  - Can talk to an API for these things, or enter manually for things like expensive work funds.

## Data Model

- Expense
  - Living Expense, Groceries, Entertainment, Restaurants, etc
- Account
  - Retirement, Checking, Savings, etc
- Contribution
  - To an `Account`
  - Retirement, Savings, etc
- Income
  - (Income is a Checking `Contribution`)

```go
type ExpenseCategory struct {
  ExpenseCategoryUUID uuid.UUID
  Name string
  Description string
}

type Expense struct {
  ExpenseUUID uuid.UUID
  Category ExpenseCategory
  Name string
  Description string
  Date time.Time
  Amount decimal.Decimal
}

type AccountCategory struct {
  AccountCategoryUUID uuid.UUID
  Name string
  Description string
}

type Account struct {
  AccountUUID uuid.UUID
  Category AccountCategory
  Name string
  Description string
  Amount decimal.Decimal
}

type Contribution struct {
  ContributionUUID uuid.UUID
  Account account.Account
  Name string
  Description string
  Date time.Time
  Amount decimal.Decimal
}
```
