import React from "react";
import api from "../common/api"
import ExpenseForm from "./ExpenseForm";
import ExpenseRow from "./ExpenseRow";

class ExpensesPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expenses: []
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  handleCreate(values) {
    api.createExpense(values)
      .then(() => this.setExpenses())
  }

  handleDelete(uuid) {
    api.deleteExpense(uuid)
      .then(() => this.setExpenses())
  }

  handleUpdate(values) {
    api.updateExpense(values)
      .then(() => this.setExpenses())
  }

  componentDidMount() {
    this.setExpenses()
  }

  setExpenses() {
    api.getExpenses()
      .then(response => {
        var expenses = (response.data === null || response.data === undefined)
          ? []
          : response.data
        this.setState({ expenses: expenses })
      })
  }

  render() {
    return (
      <div>
        <h1>Expenses</h1>
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Category</th>
              <th>Description</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {this.state.expenses.length > 0 ? (
              this.state.expenses.map(expense => (
                (
                  <ExpenseRow
                    key={expense.uuid}
                    expense={expense}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                )
              ))
            ) : (
              <tr>
                <td colSpan={6}>No Expenses</td>
              </tr>
            )}
          </tbody>
        </table>
        <ExpenseForm
          isEditMode={false}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default ExpensesPage;