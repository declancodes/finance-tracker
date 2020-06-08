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
      (response.data === null || response.data === undefined)
        ? []
        : sortBy(response.data, order)
    );
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
    return sort(
      get(ACCOUNTS_URL, filterParams),
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
    return sort(
      get(CONTRIBUTIONS_URL, filterParams),
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
    return sort(
      get(EXPENSES_URL, filterParams),
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

  deleteFund(uuid) {
    return remove(`${FUNDS_URL}/${uuid}`)
  },

  createHolding(values) {
    return create(HOLDINGS_URL, values);
  },

  getHoldings(filterParams) {
    return sort(
      get(HOLDINGS_URL, filterParams),
      ['account.name', 'fund.tickerSymbol']
    );
  },

  updateHolding(values) {
    return update(`${HOLDINGS_URL}/${values.uuid}`, values)
  },

  deleteHolding(uuid) {
    return remove(`${HOLDINGS_URL}/${uuid}`)
  }
};
