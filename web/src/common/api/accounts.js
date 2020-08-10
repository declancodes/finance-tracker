import {
  create,
  get,
  update,
  remove,
  sort,
  sortTotal
} from './base'

const ACCOUNT_CATEGORIES_URL = `${API_URL}/accountcategories`;
const ACCOUNTS_URL = `${API_URL}/accounts`;
const CONTRIBUTIONS_URL = `${API_URL}/contributions`;

export const createAccountCategory = (values) => {
  return create(ACCOUNT_CATEGORIES_URL, values);
};

export const getAccountCategories = (filterParams) => {
  return sort(
    get(ACCOUNT_CATEGORIES_URL),
    ['name']
  );
};

export const updateAccountCategory = (values) => {
  return update(`${ACCOUNT_CATEGORIES_URL}/${values.uuid}`, values)
};

export const deleteAccountCategory = (uuid) => {
  return remove(`${ACCOUNT_CATEGORIES_URL}/${uuid}`)
};

export const createAccount = (values) => {
  return create(ACCOUNTS_URL, values);
};

export const getAccounts = (filterParams) => {
  return getAccountsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
};

export const getAccountsTotal = (filterParams) => {
  return sortTotal(
    get(ACCOUNTS_URL, filterParams),
    'accounts',
    ['category.name', 'name']
  );
};

export const updateAccount = (values) => {
  return update(`${ACCOUNTS_URL}/${values.uuid}`, values)
};

export const deleteAccount = (uuid) => {
  return remove(`${ACCOUNTS_URL}/${uuid}`)
};

export const createContribution = (values) => {
  return create(CONTRIBUTIONS_URL, values);
};

export const getContributions = (filterParams) => {
  return getContributionsTotal(filterParams)
    .then(response => {
      return response.entities;
    });
};

export const getContributionsTotal = (filterParams) => {
  return sortTotal(
    get(CONTRIBUTIONS_URL, filterParams),
    'contributions',
    ['date', 'account.name', 'amount']
  );
};

export const updateContribution = (values) => {
  return update(`${CONTRIBUTIONS_URL}/${values.uuid}`, values)
};

export const deleteContribution = (uuid) => {
  return remove(`${CONTRIBUTIONS_URL}/${uuid}`)
};
