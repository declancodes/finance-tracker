export interface BaseEntity {
  uuid: string,
  name: string,
  description: string
}

export interface Total {
  total: number
}

export interface TotalEntity extends Total {
  entities: any[]
}

export interface Category extends BaseEntity {}

export interface BaseAmountEntity extends BaseEntity {
  category: string,
  amount: number
}

export interface Account extends BaseAmountEntity {}
export interface Expense extends BaseAmountEntity {
  date: Date
}

export interface Contribution extends BaseEntity {
  account: Account,
  date: Date,
  amount: number
}

export interface Fund extends BaseEntity {
  category: string,
  tickerSymbol: string,
  sharePrice: number,
  expenseRatio: number
}

export interface Holding extends BaseEntity {
  account: Account,
  fund: Fund,
  shares: number,
  value: number,
  effectiveExpense: number
}

export interface Portfolio extends BaseEntity {
  holdings: Holding[],
  assetAllocation: Map<Category, number>
}

export interface PortfolioHoldingMapping {
  uuid: string,
  portfolio: Portfolio,
  holding: Holding
}

export interface PortfolioAssetCategoryMapping {
  uuid: string,
  portfolio: Portfolio,
  assetCategory: Category,
  percentage: number
}

export interface AccountsTotal extends Total {
  accounts: Account[]
}

export interface ContributionsTotal extends Total {
  contributions: Contribution[]
}

export interface ExpensesTotal extends Total {
  expenses: Expense[]
}

export interface HoldingsTotal {
  holdings: Holding[],
  valueTotal: number,
  effectiveExpenseTotal: number
}
