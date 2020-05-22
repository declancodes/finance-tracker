import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';
import moment from 'moment';

const createContribution = (values) =>
  api.createContribution(values);

const getContributions = (start, end) =>
  api.getContributions(start, end)
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.date.localeCompare(b.date))
    );

const updateContribution = (values) =>
  api.updateContribution(values);

const deleteContribution = (uuid) =>
  api.deleteContribution(uuid);

const getOptions = () =>
  api.getAccounts()
    .then(response =>
      (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name))
    );

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
    createEntity={createContribution}
    getEntities={getContributions}
    updateEntity={updateContribution}
    deleteEntity={deleteContribution}
    getOptions={getOptions}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
