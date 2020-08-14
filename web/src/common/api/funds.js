import axios from 'axios';
import {
  create,
  get,
  update,
  remove,
  sort,
  sortTotal
} from './base'

const ASSET_CATEGORIES_URL = `${API_URL}/assetcategories`;
const FUNDS_URL = `${API_URL}/funds`;
const HOLDINGS_URL = `${API_URL}/holdings`;

export const createAssetCategory = (values) => {
  return create(ASSET_CATEGORIES_URL, values);
};

export const getAssetCategories = (filterParams) => {
  return sort(
    get(ASSET_CATEGORIES_URL),
    ['name']
  );
};

export const updateAssetCategory = (values) => {
  return update(`${ASSET_CATEGORIES_URL}/${values.uuid}`, values)
};

export const deleteAssetCategory = (uuid) => {
  return remove(`${ASSET_CATEGORIES_URL}/${uuid}`)
};

export const createFund = (values) => {
  return create(FUNDS_URL, values);
};

export const getFunds = (filterParams) => {
  return sort(
    get(FUNDS_URL, filterParams),
    ['category.name', 'name', 'tickerSymbol']
  );
};

export const updateFund = (values) => {
  return update(`${FUNDS_URL}/${values.uuid}`, values)
};

export const updateFundSharePrices = () => {
  return axios.put(FUNDS_URL);
};

export const deleteFund = (uuid) => {
  return remove(`${FUNDS_URL}/${uuid}`)
};

export const createHolding = (values) => {
  return create(HOLDINGS_URL, values);
};

export const getHoldings = (filterParams) => {
  return getHoldingsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
};

export const getHoldingsTotal = (filterParams) => {
  return sortTotal(
    get(HOLDINGS_URL, filterParams),
    'holdings',
    ['account.name', 'fund.tickerSymbol']
  );
};

export const updateHolding = (values) => {
  return update(`${HOLDINGS_URL}/${values.uuid}`, values)
};

export const deleteHolding = (uuid) => {
  return remove(`${HOLDINGS_URL}/${uuid}`)
};
