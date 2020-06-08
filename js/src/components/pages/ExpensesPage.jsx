import React from 'react';
import EntityPage from '../common/EntityPage';
import { api } from '../../common/api';
import { helpers } from '../../common/helpers';

const doExtraModifications = (values) => {
  const ecUuid = values.category;
  values.category = {
    uuid: ecUuid
  };

  const dateToSubmit = helpers.consumeDate(values.date);
  values.date = dateToSubmit;
}

const getInitialValues = (expense) => {
  let initialValues = JSON.parse(JSON.stringify(expense));
  initialValues.category = expense.category.uuid;
  initialValues.date = helpers.displayDate(expense.date);

  return initialValues;
}

export const ExpensesPage = () => (
  <EntityPage
    entityName='Expense'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      date: '',
      amount: 0
    }}
    usesFilters={true}
    usesDates={true}
    filterCategories={[
      {name: 'category', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={api.createExpense}
    getEntities={api.getExpenses}
    updateEntity={api.updateExpense}
    deleteEntity={api.deleteExpense}
    getOptions={[
      {name: 'category', value: api.getExpenseCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
