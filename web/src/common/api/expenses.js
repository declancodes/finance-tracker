import {
  create,
  get,
  update,
  remove,
  sort,
  sortTotal
} from './base'

const EXPENSE_CATEGORIES_URL = `${API_URL}/expensecategories`;
const EXPENSES_URL = `${API_URL}/expenses`;

export const createExpenseCategory = (values) => {
  return create(EXPENSE_CATEGORIES_URL, values);
};

export const getExpenseCategories = (filterParams) => {
  return sort(
    get(EXPENSE_CATEGORIES_URL),
    ['name']
  );
};

export const updateExpenseCategory = (values) => {
  return update(`${EXPENSE_CATEGORIES_URL}/${values.uuid}`, values)
};

export const deleteExpenseCategory = (uuid) => {
  return remove(`${EXPENSE_CATEGORIES_URL}/${uuid}`)
};

export const createExpense = (values) => {
  return create(EXPENSES_URL, values);
};

export const getExpenses = (filterParams) => {
  return getExpensesTotal(filterParams)
    .then(response => {
      return response.entities;
    });
};

export const getExpensesTotal = (filterParams) => {
  return sortTotal(
    get(EXPENSES_URL, filterParams),
    'expenses',
    ['date', 'category.name', 'amount']
  );
};

export const updateExpense = (values) => {
  return update(`${EXPENSES_URL}/${values.uuid}`, values)
};

export const deleteExpense = (uuid) => {
  return remove(`${EXPENSES_URL}/${uuid}`)
};
