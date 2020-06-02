import axios from 'axios';
import querystring from 'query-string';
import sortBy from 'lodash.sortby';

const API_URL = 'http://localhost:8080';
const ACCOUNT_CATEGORIES_URL = `${API_URL}/accountcategories`;
const ACCOUNTS_URL = `${API_URL}/accounts`;
const CONTRIBUTIONS_URL = `${API_URL}/contributions`;
const EXPENSE_CATEGORIES_URL = `${API_URL}/expensecategories`;
const EXPENSES_URL = `${API_URL}/expenses`;
const FUNDS_URL = `${API_URL}/funds`;
const HOLDINGS_URL = `${API_URL}/holdings`;

function create(url, values) {
  return axios.post(url, values)
}

function get(baseUrl, start, end, category, account, fund) {
  const params = {
    start: start,
    end: end,
    category: category,
    account: account,
    fund: fund
  };

  const url = querystring.stringifyUrl(
    { url: baseUrl, query: params },
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

  getAccountCategories() {
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

  getAccounts(start, end, category) {
    return sort(
      get(ACCOUNTS_URL, null, null, category),
      ['category.name', 'name']
    );
  },

  updateAccount(values) {
    return update(`${ACCOUNTS_URL}/${values.uuid}`, values)
  },

  deleteAccount(uuid) {
    return remove(`${ACCOUNTS_URL}/${uuid}`)
  },

  createContribution(values) {
    return create(CONTRIBUTIONS_URL, values);
  },

  getContributions(start, end, account) {
    return sort(
      get(CONTRIBUTIONS_URL, start, end, null, account),
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

  getExpenseCategories() {
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

  getExpenses(start, end, category) {
    return sort(
      get(EXPENSES_URL, start, end, category),
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

  getFunds() {
    return sort(
      get(FUNDS_URL),
      ['name', 'tickerSymbol']
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

  getHoldings(fundTicker) {
    return sort(
      get(HOLDINGS_URL, null, null, null, null, fundTicker),
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
