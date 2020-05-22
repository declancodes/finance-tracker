import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api'
import moment from 'moment';

class ExpensesPage extends React.Component {
  constructor(props) {
    super(props);
    this.createExpense = this.createExpense.bind(this);
    this.getExpenses = this.getExpenses.bind(this);
    this.updateExpense = this.updateExpense.bind(this);
    this.deleteExpense = this.deleteExpense.bind(this);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.getInitialValues = this.getInitialValues.bind(this);
  }

  createExpense(values) {
    return api.createExpense(values);
  }

  getExpenses(start, end) {
    return api.getExpenses(start, end)
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.date.localeCompare(b.date));
      });
  }

  updateExpense(values) {
    return api.updateExpense(values);
  }

  deleteExpense(uuid) {
    return api.deleteExpense(uuid);
  }

  getOptions() {
    return api.getExpenseCategories()
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
      });
  }

  doExtraModifications(values) {
    const ecUuid = values.category;
    values.category = {
      uuid: ecUuid
    };

    const dateToSubmit = moment(values.date).toISOString();
    values.date = dateToSubmit;
  }

  getInitialValues(expense) {
    let initialValues = JSON.parse(JSON.stringify(expense));
    initialValues.category = expense.category.uuid;
    initialValues.date = moment(expense.date).format('MM/DD/YYYY');

    return initialValues;
  }

  render() {
    return (
      <EntityPage
        entityName='Expense'
        entityPlural='Expenses'
        columnLength={6}
        blankEntity={{
          uuid: '',
          name: '',
          category: '',
          description: '',
          date: '',
          amount: 0
        }}
        usesDates={true}
        createEntity={this.createExpense}
        getEntities={this.getExpenses}
        updateEntity={this.updateExpense}
        deleteEntity={this.deleteExpense}
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        getInitialValues={this.getInitialValues}
      />
    );
  }
}

export default ExpensesPage;
