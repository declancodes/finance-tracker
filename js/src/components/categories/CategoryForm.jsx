import React from 'react';
import { EntityFormik } from '../common/forms/EntityFormik';

export const CategoryForm = ({ category, categoryType, doSubmit}) => {
  const isCreateMode = typeof category === 'undefined';
  const initialCategoryValues = {
    uuid: isCreateMode ? '' : category.uuid,
    name: isCreateMode ? '' : category.name,
    description: isCreateMode ? '' : category.description
  };
  return (
    <EntityFormik
      entityName={`${categoryType} Category`}
      entity={initialCategoryValues}
      isCreateMode={isCreateMode}
      doSubmit={doSubmit}
    />
  );
};
