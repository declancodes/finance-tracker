import React from 'react';
import EntityPage from '../EntityPage';
import {
  createContribution,
  getContributionsTotal,
  updateContribution,
  deleteContribution,
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

const getInitialValues = (contribution) => {
  let initialValues = JSON.parse(JSON.stringify(contribution));
  initialValues.account = contribution.account.uuid;
  initialValues.date = displayDate(contribution.date);

  return initialValues;
}

export const ContributionsPage = () => (
  <EntityPage
    entityName='Contribution'
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
    createEntity={createContribution}
    getEntities={getContributionsTotal}
    updateEntity={updateContribution}
    deleteEntity={deleteContribution}
    getOptions={[
      {name: 'account', value: getAccounts}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
