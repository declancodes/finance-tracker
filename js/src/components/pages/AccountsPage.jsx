import React from 'react';
import EntityPage from '../common/EntityPage';
import { api } from '../../common/api';


const doExtraModifications = (values) => {
  const acUuid = values.category;
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
    entityPlural='Accounts'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      amount: 0
    }}
    usesFilters={false}
    createEntity={api.createAccount}
    getEntities={api.getAccounts}
    updateEntity={api.updateAccount}
    deleteEntity={api.deleteAccount}
    getOptions={api.getAccountCategories}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
