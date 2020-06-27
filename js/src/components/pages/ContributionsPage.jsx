import React from 'react';
import EntityPage from './EntityPage';
import { api } from '../../common/api';
import { helpers } from '../../common/helpers';

const doExtraModifications = (values) => {
  const aUuid = values.account.value === undefined ?
    values.account :
    values.account.value;
  values.account = {
    uuid: aUuid
  };

  const dateToSubmit = helpers.consumeDate(values.date);
  values.date = dateToSubmit;
}

const getInitialValues = (contribution) => {
  let initialValues = JSON.parse(JSON.stringify(contribution));
  initialValues.account = contribution.account.uuid;
  initialValues.date = helpers.displayDate(contribution.date);

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
    usesFilters={true}
    usesDates={true}
    filterCategories={[
      {name: 'account', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={api.createContribution}
    getEntities={api.getContributions}
    updateEntity={api.updateContribution}
    deleteEntity={api.deleteContribution}
    getOptions={[
      {name: 'account', value: api.getAccounts}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
