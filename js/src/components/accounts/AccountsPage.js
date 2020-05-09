import React from "react";
import axios from "axios";
import AccountRow from "./AccountRow";
import CreateAccountForm from "./CreateAccountForm";

const API_URL = "http://localhost:8080/accounts"

class AccountsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
  }

  handleCreate(values) {
    axios.post(API_URL, values)
      .then(response => {
        console.log(response.data);
        return axios.get(API_URL);
      })
      .then(response => {
        this.setState({ accounts: response.data })
      })
  }

  handleDelete(uuid) {
    const url = `${API_URL}/${uuid}`

    axios.delete(url)
      .then(() => axios.get(API_URL))
      .then(response => {
        this.setState({ accounts: response.data })
      })
  }

  handleUpdate(values) {
    const url = `${API_URL}/${values.uuid}`

    axios.put(url, values)
      .then(response => {
        console.log(response.data);
        return axios.get(API_URL);
      })
      .then(response => {
        this.setState({ accounts: response.data })
      })
  }

  componentDidMount() {
    axios.get(API_URL).then(response => response.data)
      .then((data) => {
        this.setState({ accounts: data })
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
                    handleDelete={() => this.handleDelete(account.uuid)}
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
        <CreateAccountForm doSubmit={this.handleCreate}/>
      </div>
    );
  }
}

export default AccountsPage;