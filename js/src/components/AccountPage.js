import React from "react";
import axios from "axios";
import Account from "./Account";
import CreateAccountCategoryForm from "./CreateAccountCategoryForm";

const API_URL = "http://localhost:8080"
const ACCOUNTS_URL = `${API_URL}/accounts`

class AccountPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
  }

  componentDidMount() {
    axios.get(ACCOUNTS_URL).then(response => response.data)
    .then((data) => {
      this.setState({ accounts: data })
    })
  }

  render() {
    return (
      <div>
        <div className="accounts">
          {this.state.accounts.map(account =>
            (<Account account={account} key={account.accountUuid}/>)
          )}
        </div>
        <CreateAccountCategoryForm />
      </div>
    );
  }
}

export default AccountPage;