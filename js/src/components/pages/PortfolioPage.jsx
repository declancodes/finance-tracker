import React from 'react';
import EntityPage from './EntityPage';
import { api } from '../../common/api';

const doExtraModifications = (values) => {
  const acUuid = values.category.value;
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
    entityName='Fund'
    blankEntity={{
      uuid: '',
      name: '',
      description: '',
      holdings: '',
      assetAllocation: ''
    }}
    createEntity={api.createPortfolio}
    getEntities={api.getPortfolios}
    updateEntity={api.updatePortfolio}
    deleteEntity={api.deletePortfolio}
    getOptions={[
      {name: 'category', value: api.getAssetCategories},
      {name: 'funds', value: api.getFunds}
    ]}
    doExtraModifications={doExtraModifications}
    getInitialValues={getInitialValues}
  />
);
