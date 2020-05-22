import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';

class AccountsPage extends React.Component {
  constructor(props) {
    super(props);
    this.createAccount = this.createAccount.bind(this);
    this.getAccounts = this.getAccounts.bind(this);
    this.updateAccount = this.updateAccount.bind(this);
    this.deleteAccount = this.deleteAccount.bind(this);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.getInitialValues = this.getInitialValues.bind(this);
  }

  createAccount(values) {
    return api.createAccount(values);
  }

  getAccounts() {
    return api.getAccounts()
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.category.name.localeCompare(b.category.name));
      });
  }

  updateAccount(values) {
    return api.updateAccount(values);
  }

  deleteAccount(uuid) {
    return api.deleteAccount(uuid);
  }

  getOptions() {
    return api.getAccountCategories()
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
      });
  }

  doExtraModifications(values) {
    const acUuid = values.category;
    values.category = {
      uuid: acUuid
    };
  }

  getInitialValues(account) {
    let initialValues = JSON.parse(JSON.stringify(account));
    initialValues.category = account.category.uuid;

    return initialValues;
  }

  render() {
    return (
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
        createEntity={this.createAccount}
        getEntities={this.getAccounts}
        updateEntity={this.updateAccount}
        deleteEntity={this.deleteAccount}
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        getInitialValues={this.getInitialValues}
      />
    );
  }
}

export default AccountsPage;
