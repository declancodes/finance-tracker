import React from 'react';
import EntityPage from '../EntityPage';
import {
  createPortfolio,
  getPortfolios,
  updatePortfolio,
  deletePortfolio,
  getAssetCategories,
  getHoldings
} from '../../../common/api';

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

export const PortfolioPage = () => (
  <EntityPage
    entityName='Portfolio'
    blankEntity={{
      uuid: '',
      name: '',
      description: '',
      holdings: '',
      assetAllocation: ''
    }}
    createEntity={createPortfolio}
    getEntities={getPortfolios}
    updateEntity={updatePortfolio}
    deleteEntity={deletePortfolio}
    getOptions={[
      {name: 'category', value: getAssetCategories},
      {name: 'holdings', value: getHoldings}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
