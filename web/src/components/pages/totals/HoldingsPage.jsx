import React from 'react';
import EntityPage from '../EntityPage';
import {
  createHolding,
  getHoldingsTotal,
  updateHolding,
  deleteHolding,
  getFunds
} from '../../../common/api/funds';
import { getAccounts } from '../../../common/api/accounts';

const doExtraModifications = (values) => {
  const aUuid = values.account.value === undefined ?
    values.account :
    values.account.value;
  values.account = {
    uuid: aUuid
  };

  const fUuid = values.fund.value === undefined ?
    values.fund :
    values.fund.value;
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
      value: 0,
      effectiveExpense: 0
    }}
    hasTotal
    usesFilters
    filterCategories={[
      {name: 'account', value: '', optionValue: 'name', optionDisplay: 'name'},
      {name: 'fund', value: '', optionValue: 'tickerSymbol', optionDisplay: 'tickerSymbol'}
    ]}
    createEntity={createHolding}
    getEntities={getHoldingsTotal}
    updateEntity={updateHolding}
    deleteEntity={deleteHolding}
    getOptions={[
      {name: 'account', value: getAccounts},
      {name: 'fund', value: getFunds}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
