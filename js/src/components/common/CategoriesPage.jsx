import React from 'react';
import EntityPage from './EntityPage';

export const CategoriesPage = ({
  categoryType,
  createCategory,
  getCategories,
  updateCategory,
  deleteCategory
}) => (
  <EntityPage
    entityName={`${categoryType} Category`}
    entityPlural={`${categoryType} Categories`}
    blankEntity={{
      uuid: '',
      name: '',
      description: ''
    }}
    usesFilters={false}
    createEntity={createCategory}
    getEntities={getCategories}
    updateEntity={updateCategory}
    deleteEntity={deleteCategory}
  />
);
