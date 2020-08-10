import React from 'react';
import { CategoriesPage } from './CategoriesPage';
import {
  createAssetCategory,
  getAssetCategories,
  updateAssetCategory,
  deleteAssetCategory
} from '../../../common/api/funds';

export const AssetCategoriesPage = () => (
  <CategoriesPage
    categoryType='Asset'
    createCategory={createAssetCategory}
    getCategories={getAssetCategories}
    updateCategory={updateAssetCategory}
    deleteCategory={deleteAssetCategory}
  />
);
