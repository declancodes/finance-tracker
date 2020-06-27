import React from 'react';
import EntityPage from '../EntityPage';

export const CategoriesPage = ({
  categoryType,
  createCategory,
  getCategories,
  updateCategory,
  deleteCategory
}) => (
  <EntityPage
    entityName={`${categoryType} Category`}
    blankEntity={{
      uuid: '',
      name: '',
      description: ''
    }}
    createEntity={createCategory}
    getEntities={getCategories}
    updateEntity={updateCategory}
    deleteEntity={deleteCategory}
  />
);
