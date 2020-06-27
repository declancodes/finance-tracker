import React from 'react';
import EntityPage from './EntityPage';
import { api } from '../../common/api';

const doExtraModifications = (values) => {
  const acUuid = values.category.value === undefined ?
    values.category :
    values.category.value;
  values.category = {
    uuid: acUuid
  };
};

const getInitialValues = (account) => {
  let initialValues = JSON.parse(JSON.stringify(account));
  initialValues.category = account.category.uuid;

  return initialValues;
};

export const AccountsPage = () => (
  <EntityPage
    entityName='Account'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      amount: 0
    }}
    usesFilters
    filterCategories={[
      {name: 'category', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={api.createAccount}
    getEntities={api.getAccounts}
    updateEntity={api.updateAccount}
    deleteEntity={api.deleteAccount}
    getOptions={[
      {name: 'category', value: api.getAccountCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
