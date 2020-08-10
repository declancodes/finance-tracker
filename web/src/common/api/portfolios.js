import {
  create,
  get,
  update,
  remove,
  sort
} from './base'

const PORTFOLIOS_URL = `${API_URL}/portfolios`;
const PORTFOLIO_HOLDING_MAPPINGS_URL = `${API_URL}/portfolioholdingmappings`;
const PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL = `${API_URL}/portfolioassetcategorymappings`;

export const createPortfolio = (values) => {
  return create(PORTFOLIOS_URL, values);
};

export const getPortfolios = (filterParams) => {
  return sort(
    get(PORTFOLIOS_URL, filterParams),
    ['name']
  );
};

export const updatePortfolio = (values) => {
  return update(`${PORTFOLIOS_URL}/${values.uuid}`, values)
};

export const deletePortfolio = (uuid) => {
  return remove(`${PORTFOLIOS_URL}/${uuid}`)
};

export const createPortfolioHoldingMapping = (values) => {
  return create(PORTFOLIO_HOLDING_MAPPINGS_URL, values);
};

export const getPortfolioHoldingMappings = (filterParams) => {
  return sort(
    get(PORTFOLIO_HOLDING_MAPPINGS_URL, filterParams),
    ['uuid']
  );
};

export const updatePortfolioHoldingMapping = (values) => {
  return update(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${values.uuid}`, values)
};

export const deletePortfolioHoldingMapping = (uuid) => {
  return remove(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${uuid}`)
};

export const createPortfolioAssetCategoryMapping = (values) => {
  return create(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, values);
};

export const getPortfolioAssetCategoryMappings = (filterParams) => {
  return sort(
    get(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, filterParams),
    ['uuid']
  );
};

export const updatePortfolioAssetCategoryMapping = (values) => {
  return update(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${values.uuid}`, values)
};

export const deletePortfolioAssetCategoryMapping = (uuid) => {
  return remove(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${uuid}`)
};
