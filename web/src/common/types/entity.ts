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
  category: Category,
  amount: number
}

export interface Account extends BaseAmountEntity {
  isArchived: boolean
}

export interface Expense extends BaseAmountEntity {
  date: Date
}

export interface Contribution extends BaseEntity {
  account: Account,
  date: Date,
  amount: number
}

export interface Income extends BaseEntity {
  account: Account,
  date: Date,
  amount: number
}

export interface Fund extends BaseEntity {
  category: Category,
  tickerSymbol: string,
  sharePrice: number,
  expenseRatio: number,
  isPrivate: boolean
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

export interface IncomesTotal extends Total {
  incomes: Income[]
}

export interface ExpensesTotal extends Total {
  expenses: Expense[]
}

export interface HoldingsTotal {
  holdings: Holding[],
  valueTotal: number,
  effectiveExpenseTotal: number
}
