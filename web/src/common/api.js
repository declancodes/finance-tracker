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

const create = (url, values) => {
  return axios.post(url, values)
}

const get = (baseUrl, filterParams) => {
  const parsedUrl = {
    url: baseUrl,
    query: filterParams
  };
  const options = {
    skipNull: true,
    skipEmptyString: true
  };
  const url = querystring.stringifyUrl(parsedUrl, options);

  return axios.get(url);
}

const update = (url, values) => {
  return axios.put(url, values)
}

const remove = (url) => {
  return axios.delete(url)
}

const sort = (promise, order) => {
  return promise
    .then(response =>
      response.data === undefined || response.data === null
          ? []
          : sortBy(response.data, order)
    );
}

const sortTotal = (promise, property, order) => {
  return promise
    .then(response => {
      const hasNoData = response.data === undefined || response.data === null;
      return {
        entities: hasNoData ? [] : sortBy(response.data[property], order),
        total: hasNoData ? 0 : response.data.total
      };
    });
}

export const createAccountCategory = (values) => {
  return create(ACCOUNT_CATEGORIES_URL, values);
}

export const getAccountCategories = (filterParams) => {
  return sort(
    get(ACCOUNT_CATEGORIES_URL),
    ['name']
  );
}

export const updateAccountCategory = (values) => {
  return update(`${ACCOUNT_CATEGORIES_URL}/${values.uuid}`, values)
}

export const deleteAccountCategory = (uuid) => {
  return remove(`${ACCOUNT_CATEGORIES_URL}/${uuid}`)
}

export const createAccount = (values) => {
  return create(ACCOUNTS_URL, values);
}

export const getAccounts = (filterParams) => {
  return getAccountsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
}

export const getAccountsTotal = (filterParams) => {
  return sortTotal(
    get(ACCOUNTS_URL, filterParams),
    'accounts',
    ['category.name', 'name']
  );
}

export const updateAccount = (values) => {
  return update(`${ACCOUNTS_URL}/${values.uuid}`, values)
}

export const deleteAccount = (uuid) => {
  return remove(`${ACCOUNTS_URL}/${uuid}`)
}

export const createAssetCategory = (values) => {
  return create(ASSET_CATEGORIES_URL, values);
}

export const getAssetCategories = (filterParams) => {
  return sort(
    get(ASSET_CATEGORIES_URL),
    ['name']
  );
}

export const updateAssetCategory = (values) => {
  return update(`${ASSET_CATEGORIES_URL}/${values.uuid}`, values)
}

export const deleteAssetCategory = (uuid) => {
  return remove(`${ASSET_CATEGORIES_URL}/${uuid}`)
}

export const createContribution = (values) => {
  return create(CONTRIBUTIONS_URL, values);
}

export const getContributions = (filterParams) => {
  return getContributionsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
}

export const getContributionsTotal = (filterParams) => {
  return sortTotal(
    get(CONTRIBUTIONS_URL, filterParams),
    'contributions',
    ['date', 'account.name', 'amount']
  );
}

export const updateContribution = (values) => {
  return update(`${CONTRIBUTIONS_URL}/${values.uuid}`, values)
}

export const deleteContribution = (uuid) => {
  return remove(`${CONTRIBUTIONS_URL}/${uuid}`)
}

export const createExpenseCategory = (values) => {
  return create(EXPENSE_CATEGORIES_URL, values);
}

export const getExpenseCategories = (filterParams) => {
  return sort(
    get(EXPENSE_CATEGORIES_URL),
    ['name']
  );
}

export const updateExpenseCategory = (values) => {
  return update(`${EXPENSE_CATEGORIES_URL}/${values.uuid}`, values)
}

export const deleteExpenseCategory = (uuid) => {
  return remove(`${EXPENSE_CATEGORIES_URL}/${uuid}`)
}

export const createExpense = (values) => {
  return create(EXPENSES_URL, values);
}

export const getExpenses = (filterParams) => {
  return getExpensesTotal(filterParams)
    .then(response => {
      return response.entities;
    });
}

export const getExpensesTotal = (filterParams) => {
  return sortTotal(
    get(EXPENSES_URL, filterParams),
    'expenses',
    ['date', 'category.name', 'amount']
  );
}

export const updateExpense = (values) => {
  return update(`${EXPENSES_URL}/${values.uuid}`, values)
}

export const deleteExpense = (uuid) => {
  return remove(`${EXPENSES_URL}/${uuid}`)
}

export const createFund = (values) => {
  return create(FUNDS_URL, values);
}

export const getFunds = (filterParams) => {
  return sort(
    get(FUNDS_URL, filterParams),
    ['category.name', 'name', 'tickerSymbol']
  );
}

export const updateFund = (values) => {
  return update(`${FUNDS_URL}/${values.uuid}`, values)
}

export const updateFundSharePrices = () => {
  return axios.put(FUNDS_URL);
}

export const deleteFund = (uuid) => {
  return remove(`${FUNDS_URL}/${uuid}`)
}

export const createHolding = (values) => {
  return create(HOLDINGS_URL, values);
}

export const getHoldings = (filterParams) => {
  return getHoldingsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
}

export const getHoldingsTotal = (filterParams) => {
  return sortTotal(
    get(HOLDINGS_URL, filterParams),
    'holdings',
    ['account.name', 'fund.tickerSymbol']
  );
}

export const updateHolding = (values) => {
  return update(`${HOLDINGS_URL}/${values.uuid}`, values)
}

export const deleteHolding = (uuid) => {
  return remove(`${HOLDINGS_URL}/${uuid}`)
}

export const createPortfolio = (values) => {
  return create(PORTFOLIOS_URL, values);
}

export const getPortfolios = (filterParams) => {
  return sort(
    get(PORTFOLIOS_URL, filterParams),
    ['name']
  );
}

export const updatePortfolio = (values) => {
  return update(`${PORTFOLIOS_URL}/${values.uuid}`, values)
}

export const deletePortfolio = (uuid) => {
  return remove(`${PORTFOLIOS_URL}/${uuid}`)
}

export const createPortfolioHoldingMapping = (values) => {
  return create(PORTFOLIO_HOLDING_MAPPINGS_URL, values);
}

export const getPortfolioHoldingMappings = (filterParams) => {
  return sort(
    get(PORTFOLIO_HOLDING_MAPPINGS_URL, filterParams),
    ['uuid']
  );
}

export const updatePortfolioHoldingMapping = (values) => {
  return update(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${values.uuid}`, values)
}

export const deletePortfolioHoldingMapping = (uuid) => {
  return remove(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${uuid}`)
}

export const createPortfolioAssetCategoryMapping = (values) => {
  return create(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, values);
}

export const getPortfolioAssetCategoryMappings = (filterParams) => {
  return sort(
    get(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, filterParams),
    ['uuid']
  );
}

export const updatePortfolioAssetCategoryMapping = (values) => {
  return update(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${values.uuid}`, values)
}

export const deletePortfolioAssetCategoryMapping = (uuid) => {
  return remove(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${uuid}`)
}
