import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api'
import moment from 'moment';

const createExpense = (values) =>
  api.createExpense(values);

const getExpenses = (start, end) =>
  api.getExpenses(start, end)
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.date.localeCompare(b.date))
    );

const updateExpense = (values) =>
  api.updateExpense(values);

const deleteExpense = (uuid) =>
  api.deleteExpense(uuid);

const getOptions = () =>
  api.getExpenseCategories()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name))
    );

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
    createEntity={createExpense}
    getEntities={getExpenses}
    updateEntity={updateExpense}
    deleteEntity={deleteExpense}
    getOptions={getOptions}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
