import React from 'react';
import { Formik } from 'formik';
import moment from 'moment';
import api from '../common/api';
import { EntityForm } from '../common/forms/EntityForm';

class ExpenseForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expenseCategories: []
    };
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
    const e = this.props.expense;
    const initialExpenseValues = {
      uuid: e ? e.uuid : '',
      name: e ? e.name : '',
      category: e ? e.category.uuid : '',
      description: e ? e.description : '',
      date: e ? moment(e.date).format('MM/DD/YYYY') : '',
      amount: e ? e.amount : 0
    };

    return (
      <div>
        <h2>
          {this.props.isEditMode ? 'Edit' : 'Create'} Expense
        </h2>
        <Formik
          initialValues={initialExpenseValues}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!this.props.isEditMode) {
              delete values.uuid;
            }

            let acUuid = values.category;
            values.category = {
              uuid: acUuid
            };

            let dateToSubmit = moment(values.date).toISOString();
            values.date = dateToSubmit;

            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <EntityForm
            entity={initialExpenseValues}
            options={this.state.expenseCategories}
            isEditMode={this.props.isEditMode}
          />
        </Formik>
      </div>
    );
  }
}

export default ExpenseForm;