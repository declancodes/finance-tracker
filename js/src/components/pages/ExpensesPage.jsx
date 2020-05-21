import React from 'react';
import { DateRangePanel } from '../common/DateRangePanel';
import { EmptyEntityRow } from '../common/tables/EmptyEntityRow';
import EntityForm from '../common/forms/EntityForm';
import { EntityHeader } from '../common/tables/EntityHeader';
import EntityRow from '../common/tables/EntityRow';
import api from '../../api'
import moment from 'moment';

class ExpensesPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expenses: [],
      start: moment().startOf('month').toDate(),
      end: moment().endOf('month').toDate()
    };
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.handleStartDateSet = this.handleStartDateSet.bind(this);
    this.handleEndDateSet = this.handleEndDateSet.bind(this);
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

  handleStartDateSet(value) {
    this.setState(
      { start: value },
      () => this.setExpenses());
  }

  handleEndDateSet(value) {
    this.setState(
      { end: value },
      () => this.setExpenses());
  }

  componentDidMount() {
    this.setExpenses()
  }

  setExpenses() {
    api.getExpenses(this.state.start.toISOString(), this.state.end.toISOString())
      .then(response => {
        let expenses = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.date.localeCompare(b.date));
        this.setState({ expenses: expenses });
      });
  }

  render() {
    const entityName = 'Expense';
    const entityPlural = `${entityName}s`;
    const blankEntity = {
      uuid: '',
      name: '',
      category: '',
      description: '',
      date: '',
      amount: 0
    };

    return (
      <div>
        <h1>{entityPlural}</h1>
        <DateRangePanel
          start={this.state.start}
          end={this.state.end}
          setStart={this.handleStartDateSet}
          setEnd={this.handleEndDateSet}
        />
        <table>
          <EntityHeader entity={blankEntity}/>
          <tbody>
            {this.state.expenses.length > 0 ? (
              this.state.expenses.map(expense => {
                let initialVals = JSON.parse(JSON.stringify(expense));
                initialVals.category = expense.category.uuid;
                initialVals.date = moment(expense.date).format('MM/DD/YYYY')

                return (
                  <EntityRow
                    key={expense.uuid}
                    entityName={entityName}
                    entity={expense}
                    initialValues={initialVals}
                    getOptions={this.getOptions}
                    doExtraModifications={this.doExtraModifications}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                );
              })
            ) : (
              <EmptyEntityRow columnLength={6} entityPlural={entityPlural}/>
            )}
          </tbody>
        </table>
        <EntityForm
          entityName={entityName}
          entity={blankEntity}
          isCreateMode={true}
          getOptions={this.getOptions}
          doExtraModifications={this.doExtraModifications}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default ExpensesPage;