import { StringifiableRecord } from "query-string";
import {
  Account,
  AccountsTotal,
  Category,
  Contribution,
  ContributionsTotal,
  Income,
  IncomesTotal
} from "../types/entity";
import { create, get, update, remove, getEntities, getTotalAmountEntity } from "./base"

declare const API_URL: string;

const ACCOUNT_CATEGORIES_URL = `${API_URL}/accountcategories`;
const ACCOUNTS_URL = `${API_URL}/accounts`;
const CONTRIBUTIONS_URL = `${API_URL}/contributions`;
const INCOMES_URL = `${API_URL}/incomes`;

export const createAccountCategory = async (category: Category) => {
  return await create(ACCOUNT_CATEGORIES_URL, category);
};

export const getAccountCategories = async (filterParams: StringifiableRecord): Promise<Category[]> => {
  return await getEntities(
    get(ACCOUNT_CATEGORIES_URL, filterParams)
  );
};

export const updateAccountCategory = async (category: Category) => {
  return await update(`${ACCOUNT_CATEGORIES_URL}/${category.uuid}`, category)
};

export const deleteAccountCategory = async (uuid: string) => {
  return await remove(`${ACCOUNT_CATEGORIES_URL}/${uuid}`)
};

export const createAccount = async (account: Account) => {
  return await create(ACCOUNTS_URL, account);
};

export const getAccounts = async (filterParams: StringifiableRecord): Promise<Account[]> => {
  const accountsTotal = await getAccountsTotal(filterParams);
  return accountsTotal.accounts;
};

export const getAccountsTotal = async (filterParams: StringifiableRecord): Promise<AccountsTotal> => {
  const totalEntity = await getTotalAmountEntity(
    get(ACCOUNTS_URL, filterParams),
    "accounts"
  );

  return {
    accounts: totalEntity.entities.sort((a: Account, b: Account) => {
      if (a.category.name > b.category.name) {
        return 1;
      }

      if (a.category.name < b.category.name) {
        return -1;
      }

      return 0;
    }),
    total: totalEntity.total
  };
};

export const updateAccount = async (account: Account) => {
  return await update(`${ACCOUNTS_URL}/${account.uuid}`, account)
};

export const deleteAccount = async (uuid: string) => {
  return await remove(`${ACCOUNTS_URL}/${uuid}`)
};

export const createContribution = async (contribution: Contribution) => {
  return await create(CONTRIBUTIONS_URL, contribution);
};

export const getContributionsTotal = async (filterParams: StringifiableRecord): Promise<ContributionsTotal> => {
  const totalEntity =  await getTotalAmountEntity(
    get(CONTRIBUTIONS_URL, filterParams),
    "contributions"
  );

  return {
    contributions: totalEntity.entities.sort((a: Contribution, b: Contribution) => {
      if (a.date > b.date) {
        return 1;
      }

      if (a.date < b.date) {
        return -1;
      }

      return 0;
    }),
    total: totalEntity.total
  }
};

export const updateContribution = async (contribution: Contribution) => {
  return await update(`${CONTRIBUTIONS_URL}/${contribution.uuid}`, contribution);
};

export const deleteContribution = async (uuid: string) => {
  return await remove(`${CONTRIBUTIONS_URL}/${uuid}`);
};

export const createIncome = async (income: Income) => {
  return await create(INCOMES_URL, income);
};

export const getIncomesTotal = async (filterParams: StringifiableRecord): Promise<IncomesTotal> => {
  const totalEntity =  await getTotalAmountEntity(
    get(INCOMES_URL, filterParams),
    "incomes"
  );

  return {
    incomes: totalEntity.entities.sort((a: Income, b: Income) => {
      if (a.date > b.date) {
        return 1;
      }

      if (a.date < b.date) {
        return -1;
      }

      return 0;
    }),
    total: totalEntity.total
  }
};

export const updateIncome = async (income: Income) => {
  return await update(`${INCOMES_URL}/${income.uuid}`, income);
};

export const deleteIncome = async (uuid: string) => {
  return await remove(`${INCOMES_URL}/${uuid}`);
};
