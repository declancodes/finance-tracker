import React from "react";
import AccountForm from "./AccountForm";
import Button from "../common/Button";

class AccountRow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isEditing: false
    };
    this.handleUpdate = this.handleUpdate.bind(this);
  }

  handleUpdate(values) {
    this.props.handleUpdate(values);
    this.setEditing(false);
  }

  setEditing(val) {
    this.setState({ isEditing: val })
  }

  render() {
    const a = this.props.account;

    return (
      <tr>
        <td>{a.name}</td>
        <td>{a.category.name}</td>
        <td>{a.description}</td>
        <td>${a.amount}</td>
        <td>
          <Button
            name="Delete"
            handleFunc={() => this.props.handleDelete(a.uuid)}
          />
          {this.state.isEditing ? (
            <div>
              <AccountForm
                isEditMode={true}
                account={a}
                doSubmit={this.handleUpdate}
              />
              <Button
                name="Cancel"
                handleFunc={() => this.setEditing(false)}
              />
            </div>
          ) : (
            <Button
              name="Edit"
              handleFunc={() => this.setEditing(true)}
            />
          )}
        </td>
      </tr>
    );
  }
}

export default AccountRow;
