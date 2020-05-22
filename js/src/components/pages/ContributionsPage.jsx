import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';
import moment from 'moment';

const doExtraModifications = (values) => {
  const aUuid = values.account;
  values.account = {
    uuid: aUuid
  };

  const dateToSubmit = moment(values.date).toISOString();
  values.date = dateToSubmit;
}

const getInitialValues = (contribution) => {
  let initialValues = JSON.parse(JSON.stringify(contribution));
  initialValues.account = contribution.account.uuid;
  initialValues.date = moment(contribution.date).format('MM/DD/YYYY')

  return initialValues;
}

export const ContributionsPage = () => (
  <EntityPage
    entityName='Contribution'
    entityPlural='Contributions'
    columnLength={6}
    blankEntity={{
      uuid: '',
      name: '',
      account: '',
      description: '',
      date: '',
      amount: 0
    }}
    usesDates={true}
    createEntity={api.createContribution}
    getEntities={api.getContributions}
    updateEntity={api.updateContribution}
    deleteEntity={api.deleteContribution}
    getOptions={api.getAccounts}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
