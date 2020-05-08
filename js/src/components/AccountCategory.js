import React from "react";

class AccountCategory extends React.Component {
  render() {
    return (
      <div className="accountCategory">
        <div className="accountCategory-body">
          <h4>{this.props.accountCategory.name}</h4>
          <h5>{this.props.accountCategory.description}</h5>
        </div>
      </div>
    );
  }
}

export default AccountCategory;