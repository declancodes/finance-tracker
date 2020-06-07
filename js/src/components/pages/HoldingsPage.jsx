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
    blankEntity={{
      uuid: '',
      account: '',
      fund: '',
      shares: 0,
      value: 0
    }}
    usesFilters={false}
    createEntity={api.createHolding}
    getEntities={api.getHoldings}
    updateEntity={api.updateHolding}
    deleteEntity={api.deleteHolding}
    getOptions={[
      {key: 'account', value: api.getAccounts},
      {key: 'fund', value: api.getFunds}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
