import React from 'react';
import { CategoriesPage } from './CategoriesPage';
import {
  createExpenseCategory,
  getExpenseCategories,
  updateExpenseCategory,
  deleteExpenseCategory
} from '../../../common/api';

export const ExpenseCategoriesPage = () => (
  <CategoriesPage
    categoryType='Expense'
    createCategory={createExpenseCategory}
    getCategories={getExpenseCategories}
    updateCategory={updateExpenseCategory}
    deleteCategory={deleteExpenseCategory}
  />
);