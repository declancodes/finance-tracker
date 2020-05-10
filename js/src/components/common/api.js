import axios from "axios";

const API_URL = "http://localhost:8080"
const ACCOUNT_CATEGORIES_URL = `${API_URL}/accountcategories`
const ACCOUNTS_URL = `${API_URL}/accounts`
const CONTRIBUTIONS_URL = `${API_URL}/contributions`
const EXPENSE_CATEGORIES_URL = `${API_URL}/expensecategories`
const EXPENSES_URL = `${API_URL}/expenses`

function create(url, values) {
  return axios.post(url, values)
}

function get(url) {
  return axios.get(url)
}

function update(url, values) {
  return axios.put(url, values)
}

function remove(url) {
  return axios.delete(url)
}

const api = {
  createAccountCategory(values) {
    return create(ACCOUNT_CATEGORIES_URL, values);
  },

  getAccountCategories() {
    return get(ACCOUNT_CATEGORIES_URL);
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

  getAccounts() {
    return get(ACCOUNTS_URL);
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

  getContributions() {
    return get(CONTRIBUTIONS_URL);
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
    return get(EXPENSE_CATEGORIES_URL);
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

  getExpenses() {
    return get(EXPENSES_URL);
  },

  updateExpense(values) {
    return update(`${EXPENSES_URL}/${values.uuid}`, values)
  },

  deleteExpense(uuid) {
    return remove(`${EXPENSES_URL}/${uuid}`)
  },
}

export default api;