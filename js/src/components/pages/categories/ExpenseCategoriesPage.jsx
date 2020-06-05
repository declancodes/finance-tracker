import React from 'react';
import { CategoriesPage } from '../../common/CategoriesPage';
import { api } from '../../../common/api';

export const ExpenseCategoriesPage = () => (
  <CategoriesPage
    categoryType='Expense'
    createCategory={api.createExpenseCategory}
    getCategories={api.getExpenseCategories}
    updateCategory={api.updateExpenseCategory}
    deleteCategory={api.deleteExpenseCategory}
  />
);