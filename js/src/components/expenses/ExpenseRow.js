import React from "react";
import moment from "moment";
import { Button } from "../common/Button";
import ExpenseForm from "./ExpenseForm";
import { ModifyRowPanel } from "../common/ModifyRowPanel";

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
        <td>{moment(e.date).format("MM/DD/YYYY")}</td>
        <td>${e.amount}</td>
        <td>
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
            <ModifyRowPanel
              handleEdit={() => this.setEditing(true)}
              handleDelete={() => this.props.handleDelete(e.uuid)}
            />
          )}
        </td>
      </tr>
    );
  }
}

export default ExpenseRow;
