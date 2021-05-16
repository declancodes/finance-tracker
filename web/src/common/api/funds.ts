import { StringifiableRecord } from 'query-string';
import { Category, Fund, Holding, HoldingsTotal } from '../types/entity';
import { create, get, update, remove, getEntities } from './base';

declare const API_URL: string;

const ASSET_CATEGORIES_URL = `${API_URL}/assetcategories`;
const FUNDS_URL = `${API_URL}/funds`;
const HOLDINGS_URL = `${API_URL}/holdings`;

export const createAssetCategory = async (category: Category) => {
  return await create(ASSET_CATEGORIES_URL, category);
};

export const getAssetCategories = async (filterParams: StringifiableRecord): Promise<Category[]> => {
  return await getEntities(
    get(ASSET_CATEGORIES_URL, filterParams)
  );
};

export const updateAssetCategory = async (category: Category) => {
  return await update(`${ASSET_CATEGORIES_URL}/${category.uuid}`, category)
};

export const deleteAssetCategory = async (uuid: string) => {
  return await remove(`${ASSET_CATEGORIES_URL}/${uuid}`)
};

export const createFund = async (fund: Fund) => {
  return await create(FUNDS_URL, fund);
};

export const getFunds = async (filterParams: StringifiableRecord): Promise<Fund[]> => {
  return await getEntities(
    get(FUNDS_URL, filterParams)
  );
};

export const updateFund = async (fund: Fund) => {
  return await update(`${FUNDS_URL}/${fund.uuid}`, fund);
};

export const updateFundSharePrices = async () => {
  return await fetch(FUNDS_URL, {
    headers: { 'Content-Type': 'application/json; charset=UTF-8' },
    method: 'PUT'
  });
};

export const deleteFund = async (uuid: string) => {
  return await remove(`${FUNDS_URL}/${uuid}`)
};

export const createHolding = async (holding: Holding) => {
  return await create(HOLDINGS_URL, holding);
};

export const getHoldings = async (filterParams: StringifiableRecord): Promise<Holding[]> => {
  const holdingsTotal = await getHoldingsTotal(filterParams);

  return holdingsTotal.holdings;
};

export const getHoldingsTotal = async (filterParams: StringifiableRecord): Promise<HoldingsTotal> => {
  const holdingsResponse = await get(HOLDINGS_URL, filterParams);
  if (!holdingsResponse || !holdingsResponse.ok) {
    return {
      holdings: [],
      valueTotal: 0,
      effectiveExpenseTotal: 0
    };
  }

  const holdingsTotal = await holdingsResponse.json();
  return {
    holdings: holdingsTotal.holdings.sort((a: Holding, b: Holding) => {
      if (a.account.name > b.account.name) {
        return 1;
      }
      if (a.account.name < b.account.name) {
        return -1;
      }
      if (a.fund.tickerSymbol > b.fund.tickerSymbol) {
        return 1;
      }
      if (a.fund.tickerSymbol < b.fund.tickerSymbol) {
        return -1;
      }
      return 0;
    }),
    valueTotal: holdingsTotal.valueTotal,
    effectiveExpenseTotal: holdingsTotal.effectiveExpenseTotal
  };
};

export const updateHolding = async (holding: Holding) => {
  return await update(`${HOLDINGS_URL}/${holding.uuid}`, holding);
};

export const deleteHolding = async (uuid: string) => {
  return await remove(`${HOLDINGS_URL}/${uuid}`);
};
