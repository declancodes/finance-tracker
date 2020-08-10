import React from 'react';
import { CategoriesPage } from './CategoriesPage';
import {
  createAccountCategory,
  getAccountCategories,
  updateAccountCategory,
  deleteAccountCategory
} from '../../../common/api';

export const AccountCategoriesPage = () => (
  <CategoriesPage
    categoryType='Account'
    createCategory={createAccountCategory}
    getCategories={getAccountCategories}
    updateCategory={updateAccountCategory}
    deleteCategory={deleteAccountCategory}
  />
);