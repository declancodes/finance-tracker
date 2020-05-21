import React from 'react';
import { EmptyEntityRow } from '../common/tables/EmptyEntityRow';
import EntityFormik from '../common/forms/EntityFormik';
import { EntityHeader } from '../common/tables/EntityHeader';
import EntityRow from '../common/tables/EntityRow';
import api from '../../api'

class AccountsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
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

  handleCreate(values) {
    api.createAccount(values)
      .then(() => this.setAccounts())
  }

  handleUpdate(values) {
    api.updateAccount(values)
      .then(() => this.setAccounts())
  }

  handleDelete(uuid) {
    api.deleteAccount(uuid)
      .then(() => this.setAccounts())
  }

  componentDidMount() {
    this.setAccounts()
  }

  setAccounts() {
    api.getAccounts()
      .then(response => {
        let accounts = (response.data === null || response.data === undefined)
          ? []
          : response.data
            .sort((a, b) => a.category.name.localeCompare(b.category.name));
        this.setState({ accounts: accounts });
      })
  }

  render() {
    const entityName = 'Account';
    const entityPlural = `${entityName}s`;
    const blankEntity = {
      uuid: '',
      name: '',
      category: '',
      description: '',
      amount: 0
    };

    return (
      <div>
        <h1>{entityPlural}</h1>
        <table>
          <EntityHeader entity={blankEntity}/>
          <tbody>
            {this.state.accounts.length > 0 ? (
              this.state.accounts.map(account => {
                let initialVals = JSON.parse(JSON.stringify(account));
                initialVals.category = account.category.uuid;

                return (
                  <EntityRow
                    key={account.uuid}
                    entityName={entityName}
                    entity={account}
                    initialValues={initialVals}
                    getOptions={this.getOptions}
                    doExtraModifications={this.doExtraModifications}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                );
              })
            ) : (
              <EmptyEntityRow columnLength={5} entityPlural={entityPlural}/>
            )}
          </tbody>
        </table>
        <EntityFormik
          entityName={entityName}
          entity={blankEntity}
          isCreateMode={true}
          getOptions={this.getOptions}
          doExtraModifications={this.doExtraModifications}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default AccountsPage;