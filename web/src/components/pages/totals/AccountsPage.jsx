import React from 'react';
import EntityPage from '../EntityPage';
import {
  createAccount,
  getAccountsTotal,
  updateAccount,
  deleteAccount,
  getAccountCategories
} from '../../../common/api/accounts';

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
    createEntity={createAccount}
    getEntities={getAccountsTotal}
    updateEntity={updateAccount}
    deleteEntity={deleteAccount}
    getOptions={[
      {name: 'category', value: getAccountCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
