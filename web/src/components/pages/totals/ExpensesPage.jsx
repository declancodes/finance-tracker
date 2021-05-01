import React from 'react';
import EntityPage from '../EntityPage/EntityPage';
import {
  createExpense,
  getExpensesTotal,
  updateExpense,
  deleteExpense,
  getExpenseCategories
} from '../../../common/api/expenses';
import {
  consumeDate,
  displayDate
} from '../../../common/helpers';

const doExtraModifications = (values) => {
  const ecUuid = values.category.value === undefined ?
    values.category :
    values.category.value;
  values.category = {
    uuid: ecUuid
  };

  const dateToSubmit = consumeDate(values.date);
  values.date = dateToSubmit;
}

const getInitialValues = (expense) => {
  let initialValues = JSON.parse(JSON.stringify(expense));
  initialValues.category = expense.category.uuid;
  initialValues.date = displayDate(expense.date);

  return initialValues;
}

export const ExpensesPage = () => (
  <EntityPage
    entityName='Expense'
    entityPluralName='Expenses'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      description: '',
      date: '',
      amount: 0
    }}
    usesFilters
    usesDates
    filterCategories={[
      {name: 'category', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={createExpense}
    getEntities={getExpensesTotal}
    updateEntity={updateExpense}
    deleteEntity={deleteExpense}
    getOptions={[
      {name: 'category', value: getExpenseCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
