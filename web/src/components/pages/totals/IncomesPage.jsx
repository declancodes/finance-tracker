import React from 'react';
import EntityPage from '../EntityPage/EntityPage';
import {
  createIncome,
  getIncomesTotal,
  updateIncome,
  deleteIncome,
  getAccounts
} from '../../../common/api/accounts';
import {
  consumeDate,
  displayDate
} from '../../../common/helpers';

const doExtraModifications = (values) => {
  const aUuid = values.account.value === undefined ?
    values.account :
    values.account.value;
  values.account = {
    uuid: aUuid
  };

  const dateToSubmit = consumeDate(values.date);
  values.date = dateToSubmit;
}

const getInitialValues = (income) => {
  let initialValues = JSON.parse(JSON.stringify(income));
  initialValues.account = income.account.uuid;
  initialValues.date = displayDate(income.date);

  return initialValues;
}

export const IncomesPage = () => (
  <EntityPage
    entityName='Income'
    entityPluralName='Incomes'
    blankEntity={{
      uuid: '',
      name: '',
      account: '',
      description: '',
      date: '',
      amount: 0
    }}
    usesFilters
    usesDates
    filterCategories={[
      {name: 'account', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={createIncome}
    getEntities={getIncomesTotal}
    updateEntity={updateIncome}
    deleteEntity={deleteIncome}
    getOptions={[
      {name: 'account', value: getAccounts}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
