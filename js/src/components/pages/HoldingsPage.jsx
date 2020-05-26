import React from 'react';
import EntityPage from '../common/EntityPage';
import { api } from '../../common/api';

const doExtraModifications = (values) => {
  const aUuid = values.account;
  values.account = {
    uuid: aUuid
  };

  const fUuid = values.fund;
  values.fund = {
    uuid: fUuid
  };
};

const getInitialValues = (holding) => {
  let initialValues = JSON.parse(JSON.stringify(holding));
  initialValues.account = holding.account.uuid;
  initialValues.fund = holding.fund.uuid;

  return initialValues;
};

export const HoldingsPage = () => (
  <EntityPage
    entityName='Holding'
    entityPlural='Holdings'
    blankEntity={{
      uuid: '',
      account: '',
      fund: '',
      shares: 0
    }}
    usesFilters={false}
    createEntity={api.createHolding}
    getEntities={api.getHoldings}
    updateEntity={api.updateHolding}
    deleteEntity={api.deleteHolding}
    getOptions1={api.getAccounts}
    getOptions2={api.getFunds}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
