import React from "react";
import api from "../common/api"
import AccountForm from "./AccountForm";
import AccountRow from "./AccountRow";

class AccountsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  handleCreate(values) {
    api.createAccount(values)
      .then(() => this.setAccounts())
  }

  handleDelete(uuid) {
    api.deleteAccount(uuid)
      .then(() => this.setAccounts())
  }

  handleUpdate(values) {
    api.updateAccount(values)
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
    return (
      <div>
        <h1>Accounts</h1>
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Category</th>
              <th>Description</th>
              <th>Amount</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {this.state.accounts.length > 0 ? (
              this.state.accounts.map(account => (
                (
                  <AccountRow
                    key={account.uuid}
                    account={account}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                )
              ))
            ) : (
              <tr>
                <td colSpan={5}>No Accounts</td>
              </tr>
            )}
          </tbody>
        </table>
        <AccountForm
          isEditMode={false}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default AccountsPage;