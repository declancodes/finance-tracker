import React from 'react';
import EntityPage from '../EntityPage';
import {
  createPortfolio,
  getPortfolios,
  updatePortfolio,
  deletePortfolio
} from '../../../common/api/portfolios';
import {
  getAssetCategories,
  getHoldings
} from '../../../common/api/funds';

const doExtraModifications = (values) => {
  const assetAllocation = values.assetAllocation
    .reduce((acc, aa) => {
      if (aa.category !== '') {
        acc.push({
          category: { uuid: aa.category },
          percentage: aa.percentage
        });
      }
      return acc;
    }, []);
  values.assetAllocation = assetAllocation;

  const holdings = values.holdings.map(h => {
    return {
      holding: { uuid: h }
    };
  });
  values.holdings = holdings;
};

const getInitialValues = (portfolio) => {
  let initialValues = JSON.parse(JSON.stringify(portfolio));
  initialValues.holdings = portfolio.holdings.map(h => h.holding.uuid);

  return initialValues;
};

export const PortfolioPage = () => (
  <EntityPage
    entityName='Portfolio'
    entityPluralName='Portfolios'
    blankEntity={{
      uuid: '',
      name: '',
      description: '',
      holdings: [],
      assetAllocation: []
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
