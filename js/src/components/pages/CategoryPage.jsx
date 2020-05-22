import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';

const createAccountCategory = (values) =>
  api.createAccountCategory(values);

const getAccountCategories = () =>
  api.getAccountCategories()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name))
    );

const updateAccountCategory = (values) =>
  api.updateAccountCategory(values);

const deleteAccountCategory = (uuid) =>
  api.deleteAccountCategory(uuid);

const createExpenseCategory = (values) =>
  api.createExpenseCategory(values);

const getExpenseCategories = () =>
  api.getExpenseCategories()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name))
    );

const updateExpenseCategory = (values) =>
  api.updateExpenseCategory(values);

const deleteExpenseCategory = (uuid) =>
  api.deleteExpenseCategory(uuid);

export const CategoryPage = ({ categoryType }) => {
  const isAccountCategory = categoryType === 'Account';
  return (
    <EntityPage
      entityName={`${categoryType} Category`}
      entityPlural={`${categoryType} Categories`}
      columnLength={3}
      blankEntity={{
        uuid: '',
        name: '',
        description: ''
      }}
      usesDates={false}
      createEntity={isAccountCategory ? createAccountCategory : createExpenseCategory}
      getEntities={isAccountCategory ? getAccountCategories : getExpenseCategories}
      updateEntity={isAccountCategory ? updateAccountCategory : updateExpenseCategory}
      deleteEntity={isAccountCategory ? deleteAccountCategory : deleteExpenseCategory}
    />
  );
};
