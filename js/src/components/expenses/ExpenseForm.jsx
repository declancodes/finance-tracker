import React from 'react';
import EntityFormik from '../common/forms/EntityFormik';
import api from '../common/api';
import moment from 'moment';

class ExpenseForm extends React.Component {
  constructor(props) {
    super(props);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
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
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default ExpenseForm;
