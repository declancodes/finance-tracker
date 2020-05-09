import React from "react";
import AccountForm from "./AccountForm";
import DeleteButton from "../common/DeleteButton";

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
          <DeleteButton handleDelete={() => this.props.handleDelete(a.uuid)}/>
          {this.state.isEditing ? (
            <div>
              <AccountForm
                isEditMode={true}
                account={a}
                doSubmit={this.handleUpdate}
              />
              <button onClick={() => this.setEditing(false)}>
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="edit-button"
              onClick={() => this.setEditing(true)}>
                Edit
            </button>
          )}
        </td>
      </tr>
    );
  }
}

export default AccountRow;
