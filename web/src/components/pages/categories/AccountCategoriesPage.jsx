import React from 'react';
import { CategoriesPage } from './CategoriesPage';
import { api } from '../../../common/api';

export const AccountCategoriesPage = () => (
  <CategoriesPage
    categoryType='Account'
    createCategory={api.createAccountCategory}
    getCategories={api.getAccountCategories}
    updateCategory={api.updateAccountCategory}
    deleteCategory={api.deleteAccountCategory}
  />
);