import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api'
import moment from 'moment';

const doExtraModifications = (values) => {
  const ecUuid = values.category;
  values.category = {
    uuid: ecUuid
  };

  const dateToSubmit = moment(values.date).toISOString();
  values.date = dateToSubmit;
}

const getInitialValues = (expense) => {
  let initialValues = JSON.parse(JSON.stringify(expense));
  initialValues.category = expense.category.uuid;
  initialValues.date = moment(expense.date).format('MM/DD/YYYY');

  return initialValues;
}

export const ExpensesPage = () => (
  <EntityPage
    entityName='Expense'
    entityPlural='Expenses'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      date: '',
      amount: 0
    }}
    usesDates={true}
    createEntity={api.createExpense}
    getEntities={api.getExpenses}
    updateEntity={api.updateExpense}
    deleteEntity={api.deleteExpense}
    getOptions={api.getExpenseCategories}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
