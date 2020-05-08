import React from "react";
import axios from "axios";
import AccountCategory from "./AccountCategory";
import CreateAccountCategoryForm from "./CreateAccountCategoryForm";
import DeleteButton from "./DeleteButton"

const API_URL = "http://localhost:8080/accountcategories"

class AccountCategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accountCategories: []
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleClick = this.handleClick.bind(this);
  }

  handleSubmit(values) {
    axios.post(API_URL, values)
    .then(response => {
      console.log(response.data);
      return axios.get(API_URL);
    })
    .then(response => {
      this.setState({ accountCategories: response.data })
    })
  }

  handleClick(uuid) {
    const url = `${API_URL}/${uuid}`

    axios.delete(url)
    .then(() => axios.get(API_URL))
    .then(response => {
      this.setState({ accountCategories: response.data })
    })
  }

  componentDidMount() {
    axios.get(API_URL).then(response => response.data)
    .then((data) => {
      this.setState({ accountCategories: data })
    })
  }

  render() {
    return (
      <div>
        <h1>Account Categories</h1>
        <div className="accountCategories">
          {this.state.accountCategories.map(accountCategory =>
            (
              <div key={"container-" + accountCategory.accountCategoryUuid}>
                <AccountCategory
                  accountCategory={accountCategory}
                  key={accountCategory.accountCategoryUuid}
                />
                <DeleteButton
                  doClick={() => this.handleClick(accountCategory.accountCategoryUuid)}
                  key={"delete-" + accountCategory.accountCategoryUuid}
                />
              </div>
            )
          )}
        </div>
        <CreateAccountCategoryForm doSubmit={this.handleSubmit}/>
      </div>
    );
  }
}

export default AccountCategoryPage;