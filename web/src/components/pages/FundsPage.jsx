import React from 'react';
import { Button } from 'react-bootstrap';
import EntityPage from './EntityPage';
import {
  createFund,
  getFunds,
  updateFund,
  deleteFund,
  getAssetCategories,
  updateFundSharePrices
} from '../../common/api/funds';

const doExtraModifications = (values) => {
  const acUuid = values.category.value === undefined ?
    values.category :
    values.category.value;
  values.category = {
    uuid: acUuid
  };
};

const getInitialValues = (fund) => {
  let initialValues = JSON.parse(JSON.stringify(fund));
  initialValues.category = fund.category.uuid;

  return initialValues;
};

export const FundsPage = () => (
  <EntityPage
    entityName='Fund'
    entityPluralName='Funds'
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      tickerSymbol: '',
      sharePrice: 0,
      expenseRatio: 0
    }}
    usesFilters
    filterCategories={[
      {name: 'category', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={createFund}
    getEntities={getFunds}
    updateEntity={updateFund}
    deleteEntity={deleteFund}
    getOptions={[
      {name: 'category', value: getAssetCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  >
    <Button onClick={updateFundSharePrices}>
      Get latest share prices
    </Button>
  </EntityPage>
);
