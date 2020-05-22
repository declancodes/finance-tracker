import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';

const createAccount = (values) =>
  api.createAccount(values);

const getAccounts = () =>
  api.getAccounts()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.category.name.localeCompare(b.category.name))
    );

const updateAccount = (values) =>
  api.updateAccount(values);

const deleteAccount = (values) =>
  api.deleteAccount(values);

const getOptions = () =>
  api.getAccountCategories()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name))
    );

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
    columnLength={5}
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      amount: 0
    }}
    usesDates={false}
    createEntity={createAccount}
    getEntities={getAccounts}
    updateEntity={updateAccount}
    deleteEntity={deleteAccount}
    getOptions={getOptions}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
