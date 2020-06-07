import React from 'react';
import EntityPage from '../common/EntityPage';
import { api } from '../../common/api';

const doExtraModifications = (values) => {
  const acUuid = values.category;
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
    usesDates={false}
    filterCategoryName='category'
    createEntity={api.createFund}
    getEntities={api.getFunds}
    updateEntity={api.updateFund}
    deleteEntity={api.deleteFund}
    getOptions={[
      {key: 'category', value: api.getAssetCategories}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
