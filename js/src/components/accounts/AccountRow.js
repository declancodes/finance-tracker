import React from "react";
import DeleteButton from "../common/DeleteButton";
import EditAccountForm from "./EditAccountForm";

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
    this.setState({ isEditing: false });
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
          <DeleteButton handleDelete={() => this.props.handleDelete()}/>
          {this.state.isEditing ? (
            <div>
              <EditAccountForm
                account={a}
                doUpdate={this.handleUpdate}
              />
              <button onClick={() => this.setState({ isEditing: false })}>
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="edit-button"
              onClick={() => this.setState({ isEditing: true })}>
                Edit
            </button>
          )}
        </td>
      </tr>
    );
  }
}

export default AccountRow;
