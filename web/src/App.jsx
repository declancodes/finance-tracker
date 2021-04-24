import React from 'react';
import { hot } from 'react-hot-loader';
import {
  BrowserRouter,
  Switch,
  Route,
} from 'react-router-dom';
import { AccountsPage } from './components/pages/totals/AccountsPage';
import { AccountCategoriesPage } from './components/pages/categories/AccountCategoriesPage';
import { AssetCategoriesPage } from './components/pages/categories/AssetCategoriesPage';
import { ContributionsPage } from './components/pages/totals/ContributionsPage';
import { ExpensesPage } from './components/pages/totals/ExpensesPage';
import { ExpenseCategoriesPage } from './components/pages/categories/ExpenseCategoriesPage';
import { FundsPage } from './components/pages/FundsPage';
import { HoldingsPage } from './components/pages/totals/HoldingsPage';
import { HomePage } from './components/pages/HomePage';
import { PortfolioPage } from './components/pages/totals/PortfolioPage';
import { Dropdown } from './components/nav/Dropdown';

const accounts = '/accounts';
const accountCategories = '/accountcategories';
const contributions = '/contributions';
const expenses = '/expenses';
const expenseCategories = '/expensecategories';
const funds = '/funds';
const assetCategories = '/assetcategories';
const holdings = '/holdings';
const portfolios = '/portfolios';

const App = () => {
  return (
    <BrowserRouter>
      <nav>
        <a href='/'>Finance Tracker</a>
        <div className='mr-auto'>
          <Dropdown
            title='Accounts'
            linkItems={[
              {link: accounts, display: 'Accounts'},
              {link: accountCategories, display: 'Account Categories'},
            ]}
          />
          <a href={contributions}>Contributions</a>
          <Dropdown
            title='Expenses'
            linkItems={[
              {link: expenses, display: 'Expenses'},
              {link: expenseCategories, display: 'Expense Categories'},
            ]}
          />
          <a href={holdings}>Holdings</a>
          <Dropdown
            title='Funds'
            linkItems={[
              {link: funds, display: 'Funds'},
              {link: assetCategories, display: 'Asset Categories'},
            ]}
          />
          <a href={portfolios}>Portfolios</a>
        </div>
      </nav>
      <Switch>
        <Route path='/' exact component={HomePage}/>
        <Route path={accounts} component={AccountsPage}/>
        <Route path={accountCategories} component={AccountCategoriesPage}/>
        <Route path={contributions} component={ContributionsPage}/>
        <Route path={expenses} component={ExpensesPage}/>
        <Route path={expenseCategories} component={ExpenseCategoriesPage}/>
        <Route path={funds} component={FundsPage}/>
        <Route path={assetCategories} component={AssetCategoriesPage}/>
        <Route path={holdings} component={HoldingsPage}/>
        <Route path={portfolios} component={PortfolioPage}/>
      </Switch>
    </BrowserRouter>
  );
};

export default hot(module)(App);