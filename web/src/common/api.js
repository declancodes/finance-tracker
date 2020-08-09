import axios from 'axios';
import querystring from 'query-string';
import sortBy from 'lodash.sortby';

const API_URL = 'http://localhost:8080';
const ACCOUNT_CATEGORIES_URL = `${API_URL}/accountcategories`;
const ACCOUNTS_URL = `${API_URL}/accounts`;
const ASSET_CATEGORIES_URL = `${API_URL}/assetcategories`;
const CONTRIBUTIONS_URL = `${API_URL}/contributions`;
const EXPENSE_CATEGORIES_URL = `${API_URL}/expensecategories`;
const EXPENSES_URL = `${API_URL}/expenses`;
const FUNDS_URL = `${API_URL}/funds`;
const HOLDINGS_URL = `${API_URL}/holdings`;
const PORTFOLIOS_URL = `${API_URL}/portfolios`;
const PORTFOLIO_HOLDING_MAPPINGS_URL = `${API_URL}/portfolioholdingmappings`;
const PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL = `${API_URL}/portfolioassetcategorymappings`;

function create(url, values) {
  return axios.post(url, values)
}

function get(baseUrl, filterParams) {
  const url = querystring.stringifyUrl(
    { url: baseUrl, query: filterParams },
    { skipNull: true, skipEmptyString: true }
  );

  return axios.get(url);
}

function update(url, values) {
  return axios.put(url, values)
}

function remove(url) {
  return axios.delete(url)
}

function sort(promise, order) {
  return promise
    .then(response =>
      response.data === undefined || response.data === null
          ? []
          : sortBy(response.data, order)
    );
}

function sortTotal(promise, property, order) {
  return promise
    .then(response => {
      const hasNoData = response.data === undefined || response.data === null;
      return {
        entities: hasNoData ? [] : sortBy(response.data[property], order),
        total: hasNoData ? 0 : response.data.total
      };
    });
}

export const api = {
  createAccountCategory(values) {
    return create(ACCOUNT_CATEGORIES_URL, values);
  },

  getAccountCategories(filterParams) {
    return sort(
      get(ACCOUNT_CATEGORIES_URL),
      ['name']
    );
  },

  updateAccountCategory(values) {
    return update(`${ACCOUNT_CATEGORIES_URL}/${values.uuid}`, values)
  },

  deleteAccountCategory(uuid) {
    return remove(`${ACCOUNT_CATEGORIES_URL}/${uuid}`)
  },

  createAccount(values) {
    return create(ACCOUNTS_URL, values);
  },

  getAccounts(filterParams) {
    return api.getAccountsTotal(filterParams)
      .then(response => {
        return response.entities;
      });
  },

  getAccountsTotal(filterParams) {
    return sortTotal(
      get(ACCOUNTS_URL, filterParams),
      'accounts',
      ['category.name', 'name']
    );
  },

  updateAccount(values) {
    return update(`${ACCOUNTS_URL}/${values.uuid}`, values)
  },

  deleteAccount(uuid) {
    return remove(`${ACCOUNTS_URL}/${uuid}`)
  },

  createAssetCategory(values) {
    return create(ASSET_CATEGORIES_URL, values);
  },

  getAssetCategories(filterParams) {
    return sort(
      get(ASSET_CATEGORIES_URL),
      ['name']
    );
  },

  updateAssetCategory(values) {
    return update(`${ASSET_CATEGORIES_URL}/${values.uuid}`, values)
  },

  deleteAssetCategory(uuid) {
    return remove(`${ASSET_CATEGORIES_URL}/${uuid}`)
  },

  createContribution(values) {
    return create(CONTRIBUTIONS_URL, values);
  },

  getContributions(filterParams) {
    return api.getContributionsTotal(filterParams)
      .then(response => {
        return response.entities;
      });
  },

  getContributionsTotal(filterParams) {
    return sortTotal(
      get(CONTRIBUTIONS_URL, filterParams),
      'contributions',
      ['date', 'account.name', 'amount']
    );
  },

  updateContribution(values) {
    return update(`${CONTRIBUTIONS_URL}/${values.uuid}`, values)
  },

  deleteContribution(uuid) {
    return remove(`${CONTRIBUTIONS_URL}/${uuid}`)
  },

  createExpenseCategory(values) {
    return create(EXPENSE_CATEGORIES_URL, values);
  },

  getExpenseCategories(filterParams) {
    return sort(
      get(EXPENSE_CATEGORIES_URL),
      ['name']
    );
  },

  updateExpenseCategory(values) {
    return update(`${EXPENSE_CATEGORIES_URL}/${values.uuid}`, values)
  },

  deleteExpenseCategory(uuid) {
    return remove(`${EXPENSE_CATEGORIES_URL}/${uuid}`)
  },

  createExpense(values) {
    return create(EXPENSES_URL, values);
  },

  getExpenses(filterParams) {
    return api.getExpensesTotal(filterParams)
      .then(response => {
        return response.entities;
      });
  },

  getExpensesTotal(filterParams) {
    return sortTotal(
      get(EXPENSES_URL, filterParams),
      'expenses',
      ['date', 'category.name', 'amount']
    );
  },

  updateExpense(values) {
    return update(`${EXPENSES_URL}/${values.uuid}`, values)
  },

  deleteExpense(uuid) {
    return remove(`${EXPENSES_URL}/${uuid}`)
  },

  createFund(values) {
    return create(FUNDS_URL, values);
  },

  getFunds(filterParams) {
    return sort(
      get(FUNDS_URL, filterParams),
      ['category.name', 'name', 'tickerSymbol']
    );
  },

  updateFund(values) {
    return update(`${FUNDS_URL}/${values.uuid}`, values)
  },

  updateFundSharePrices() {
    return axios.put(FUNDS_URL);
  },

  deleteFund(uuid) {
    return remove(`${FUNDS_URL}/${uuid}`)
  },

  createHolding(values) {
    return create(HOLDINGS_URL, values);
  },

  getHoldings(filterParams) {
    return api.getHoldingsTotal(filterParams)
      .then(response => {
        return response.entities;
      });
  },

  getHoldingsTotal(filterParams) {
    return sortTotal(
      get(HOLDINGS_URL, filterParams),
      'holdings',
      ['account.name', 'fund.tickerSymbol']
    );
  },

  updateHolding(values) {
    return update(`${HOLDINGS_URL}/${values.uuid}`, values)
  },

  deleteHolding(uuid) {
    return remove(`${HOLDINGS_URL}/${uuid}`)
  },

  createPortfolio(values) {
    return create(PORTFOLIOS_URL, values);
  },

  getPortfolios(filterParams) {
    return sort(
      get(PORTFOLIOS_URL, filterParams),
      ['name']
    );
  },

  updatePortfolio(values) {
    return update(`${PORTFOLIOS_URL}/${values.uuid}`, values)
  },

  deletePortfolio(uuid) {
    return remove(`${PORTFOLIOS_URL}/${uuid}`)
  },

  createPortfolioHoldingMapping(values) {
    return create(PORTFOLIO_HOLDING_MAPPINGS_URL, values);
  },

  getPortfolioHoldingMappings(filterParams) {
    return sort(
      get(PORTFOLIO_HOLDING_MAPPINGS_URL, filterParams),
      ['uuid']
    );
  },

  updatePortfolioHoldingMapping(values) {
    return update(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${values.uuid}`, values)
  },

  deletePortfolioHoldingMapping(uuid) {
    return remove(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${uuid}`)
  },

  createPortfolioAssetCategoryMapping(values) {
    return create(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, values);
  },

  getPortfolioAssetCategoryMappings(filterParams) {
    return sort(
      get(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, filterParams),
      ['uuid']
    );
  },

  updatePortfolioAssetCategoryMapping(values) {
    return update(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${values.uuid}`, values)
  },

  deletePortfolioAssetCategoryMapping(uuid) {
    return remove(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${uuid}`)
  }
};
