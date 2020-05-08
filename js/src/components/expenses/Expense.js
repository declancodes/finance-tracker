import React from "react";

class Expense extends React.Component {
  render() {
    const e = this.props.expense;

    return (
      <div className="expense">
        <div className="expense-body">
          <h5>{e.name}</h5>
          <h6>Type: {e.expenseCategory.name}</h6>
          <h6>{e.description}</h6>
          <h6>Amount: ${e.amount}</h6>
          <h6>Date: {e.date}</h6>
        </div>
      </div>
    );
  }
}

export default Expense;