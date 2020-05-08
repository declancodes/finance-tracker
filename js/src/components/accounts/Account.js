import React from "react";

class Account extends React.Component {
  render() {
    const a = this.props.account;

    return (
      <div className="account">
        <div className="account-body">
          <h5>{a.name}</h5>
          <h6>Type: {a.accountCategory.name}</h6>
          <h6>{a.description}</h6>
          <h6>Amount: ${a.amount}</h6>
        </div>
      </div>
    );
  }
}

export default Account;