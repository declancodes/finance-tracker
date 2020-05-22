import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';

export const CategoryPage = ({ categoryType }) => {
  const isAccountCategory = categoryType === 'Account';
  return (
    <EntityPage
      entityName={`${categoryType} Category`}
      entityPlural={`${categoryType} Categories`}
      blankEntity={{
        uuid: '',
        name: '',
        description: ''
      }}
      usesDates={false}
      createEntity={isAccountCategory ? api.createAccountCategory : api.createExpenseCategory}
      getEntities={isAccountCategory ? api.getAccountCategories : api.getExpenseCategories}
      updateEntity={isAccountCategory ? api.updateAccountCategory : api.updateExpenseCategory}
      deleteEntity={isAccountCategory ? api.deleteAccountCategory : api.deleteExpenseCategory}
    />
  );
};
