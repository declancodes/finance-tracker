import React from 'react';
import { Button } from 'react-bootstrap';
import EntityPage from './EntityPage';
import { api } from '../../common/api';

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
    blankEntity={{
      uuid: '',
      name: '',
      category: '',
      tickerSymbol: '',
      sharePrice: 0,
      expenseRatio: 0
    }}
    usesFilters={true}
    filterCategories={[
      {name: 'category', value: '', optionValue: 'name', optionDisplay: 'name'}
    ]}
    createEntity={api.createFund}
    getEntities={api.getFunds}
    updateEntity={api.updateFund}
    deleteEntity={api.deleteFund}
    getOptions={[
      {name: 'category', value: api.getAssetCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  >
    <Button onClick={api.updateFundSharePrices}>
      Get latest share prices
    </Button>
  </EntityPage>
);
