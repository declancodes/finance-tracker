import React from 'react';
import { CategoriesPage } from '../../common/CategoriesPage';
import { api } from '../../../common/api';

export const AssetCategoriesPage = () => (
  <CategoriesPage
    categoryType='Asset'
    createCategory={api.createAssetCategory}
    getCategories={api.getAssetCategories}
    updateCategory={api.updateAssetCategory}
    deleteCategory={api.deleteAssetCategory}
  />
);
