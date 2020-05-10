import React from "react";
import Button from "../common/Button";
import ExpenseForm from "./ExpenseForm";

class ExpenseRow extends React.Component {
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
    const e = this.props.expense;

    return (
      <tr>
        <td>{e.name}</td>
        <td>{e.category.name}</td>
        <td>{e.description}</td>
        <td>{e.date}</td>
        <td>${e.amount}</td>
        <td>
          <Button
            name="Delete"
            handleFunc={() => this.props.handleDelete(e.uuid)}
          />
          {this.state.isEditing ? (
            <div>
              <ExpenseForm
                isEditMode={true}
                expense={e}
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

export default ExpenseRow;
