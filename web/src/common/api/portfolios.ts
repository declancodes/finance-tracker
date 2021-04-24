import { StringifiableRecord } from 'query-string';
import { Portfolio, PortfolioAssetCategoryMapping, PortfolioHoldingMapping } from '../types/entity';
import { create, get, update, remove, getEntities } from './base'

declare const API_URL: string;

const PORTFOLIOS_URL = `${API_URL}/portfolios`;
const PORTFOLIO_HOLDING_MAPPINGS_URL = `${API_URL}/portfolioholdingmappings`;
const PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL = `${API_URL}/portfolioassetcategorymappings`;

export const createPortfolio = async (portfolio: Portfolio) => {
  return await create(PORTFOLIOS_URL, portfolio);
};

export const getPortfolios = async (filterParams: StringifiableRecord): Promise<Portfolio[]> => {
  return await getEntities(
    get(PORTFOLIOS_URL, filterParams)
  );
};

export const updatePortfolio = async (portfolio: Portfolio) => {
  return await update(`${PORTFOLIOS_URL}/${portfolio.uuid}`, portfolio);
};

export const deletePortfolio = async (uuid: string) => {
  return await remove(`${PORTFOLIOS_URL}/${uuid}`);
};

export const createPortfolioHoldingMapping = async (portfolioHoldingMapping: PortfolioHoldingMapping) => {
  return await create(PORTFOLIO_HOLDING_MAPPINGS_URL, portfolioHoldingMapping);
};

export const getPortfolioHoldingMappings = async (filterParams: StringifiableRecord): Promise<PortfolioHoldingMapping[]> => {
  return await getEntities(
    get(PORTFOLIO_HOLDING_MAPPINGS_URL, filterParams)
  );
};

export const updatePortfolioHoldingMapping = async (portfolioHoldingMapping: PortfolioHoldingMapping) => {
  return await update(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${portfolioHoldingMapping.uuid}`, portfolioHoldingMapping);
};

export const deletePortfolioHoldingMapping = async (uuid: string) => {
  return remove(`${PORTFOLIO_HOLDING_MAPPINGS_URL}/${uuid}`);
};

export const createPortfolioAssetCategoryMapping = async (portfolioAssetCategoryMapping: PortfolioAssetCategoryMapping) => {
  return await create(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, portfolioAssetCategoryMapping);
};

export const getPortfolioAssetCategoryMappings = async (filterParams: StringifiableRecord): Promise<PortfolioAssetCategoryMapping[]> => {
  return getEntities(
    get(PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL, filterParams)
  );
};

export const updatePortfolioAssetCategoryMapping = async (portfolioAssetCategoryMapping: PortfolioAssetCategoryMapping) => {
  return update(
    `${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${portfolioAssetCategoryMapping.uuid}`, portfolioAssetCategoryMapping
  );
};

export const deletePortfolioAssetCategoryMapping = async (uuid: string) => {
  return await remove(`${PORTFOLIO_ASSET_CATEGORY_MAPPINGS_URL}/${uuid}`);
};
