import React from 'react';
import { EntityFormik } from '../common/forms/EntityFormik';
import api from '../common/api';
import moment from 'moment';

class ExpenseForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expenseCategories: []
    };
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
  }

  doExtraModifications(values) {
    const acUuid = values.category;
    values.category = {
      uuid: acUuid
    };

    const dateToSubmit = moment(values.date).toISOString();
    values.date = dateToSubmit;
  }

  doSubmit(values) {
    this.props.doSubmit(values);
  }

  componentDidMount() {
    api.getExpenseCategories()
      .then(response => {
        const expenseCategories = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
        this.setState({ expenseCategories: expenseCategories })
      })
  }

  render() {
    const isCreating = this.props.isCreateMode;
    const e = this.props.expense;
    const entity = {
      uuid: isCreating ? '' : e.uuid,
      name: isCreating ? '' : e.name,
      category: isCreating ? '' : e.category.uuid,
      description: isCreating ? '' : e.description,
      date: isCreating ? '' : moment(e.date).format('MM/DD/YYYY'),
      amount: isCreating ? 0 : e.amount
    };

    return (
      <EntityFormik
        entityName='Expense'
        entity={entity}
        isCreateMode={isCreating}
        options={this.state.expenseCategories}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default ExpenseForm;
