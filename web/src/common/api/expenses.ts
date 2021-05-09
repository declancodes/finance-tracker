import { StringifiableRecord } from 'query-string';
import { Category, Expense, ExpensesTotal } from '../types/entity';
import { create, get, update, remove, getEntities, getTotalAmountEntity } from './base'

declare const API_URL: string;

const EXPENSE_CATEGORIES_URL = `${API_URL}/expensecategories`;
const EXPENSES_URL = `${API_URL}/expenses`;

export const createExpenseCategory = async (category: Category) => {
  return await create(EXPENSE_CATEGORIES_URL, category);
};

export const getExpenseCategories = async (filterParams: StringifiableRecord): Promise<Category[]> => {
  return await getEntities(
    get(EXPENSE_CATEGORIES_URL, filterParams)
  );
};

export const updateExpenseCategory = async (category: Category) => {
  return await update(`${EXPENSE_CATEGORIES_URL}/${category.uuid}`, category)
};

export const deleteExpenseCategory = async (uuid: string) => {
  return await remove(`${EXPENSE_CATEGORIES_URL}/${uuid}`)
};

export const createExpense = async (expense: Expense) => {
  return await create(EXPENSES_URL, expense);
};

export const getExpensesTotal = async (filterParams: StringifiableRecord): Promise<ExpensesTotal> => {
  const totalEntity = await getTotalAmountEntity(
    get(EXPENSES_URL, filterParams),
    'expenses'
  );

  return {
    expenses: totalEntity.entities.sort((a: Expense, b: Expense) => {
      if (a.date > b.date) {
        return 1;
      }

      if (a.date < b.date) {
        return -1;
      }

      return 0;
    }),
    total: totalEntity.total
  };
};

export const updateExpense = async (expense: Expense) => {
  return await update(`${EXPENSES_URL}/${expense.uuid}`, expense);
};

export const deleteExpense = async (uuid: string) => {
  return await remove(`${EXPENSES_URL}/${uuid}`)
};
