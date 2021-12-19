import React from 'react';
import { hot } from 'react-hot-loader';
import {
  BrowserRouter,
  Switch,
  Route,
} from 'react-router-dom';
import { NavBar } from './components/nav/NavBar/NavBar';
import { NavDropdown } from './components/nav/NavDropdown/NavDropdown';
import { NavItem } from './components/nav/NavItem/NavItem';
import { AccountsPage } from './components/pages/totals/AccountsPage';
import { AccountCategoriesPage } from './components/pages/categories/AccountCategoriesPage';
import { AssetCategoriesPage } from './components/pages/categories/AssetCategoriesPage';
import { ContributionsPage } from './components/pages/totals/ContributionsPage';
import { IncomesPage } from './components/pages/totals/IncomesPage';
import { ExpensesPage } from './components/pages/totals/ExpensesPage';
import { ExpenseCategoriesPage } from './components/pages/categories/ExpenseCategoriesPage';
import { FundsPage } from './components/pages/FundsPage';
import { HoldingsPage } from './components/pages/totals/HoldingsPage';
import { HomePage } from './components/pages/HomePage';
import { PortfolioPage } from './components/pages/totals/PortfolioPage';

const accounts = '/accounts';
const accountCategories = '/accountcategories';
const contributions = '/contributions';
const incomes = '/incomes';
const expenses = '/expenses';
const expenseCategories = '/expensecategories';
const funds = '/funds';
const assetCategories = '/assetcategories';
const holdings = '/holdings';
const portfolios = '/portfolios';

const App = () => {
  return (
    <BrowserRouter>
      <NavBar>
        <NavItem to='/'>
          Finance Tracker
        </NavItem>
        <NavDropdown
          title='Accounts'
          navItems={[
            {to: accounts, title: 'Accounts'},
            {to: accountCategories, title: 'Account Categories'},
          ]}
        />
        <NavItem to={contributions}>
          Contributions
        </NavItem>
        <NavItem to={incomes}>
          Incomes
        </NavItem>
        <NavDropdown
          title='Expenses'
          navItems={[
            {to: expenses, title: 'Expenses'},
            {to: expenseCategories, title: 'Expense Categories'},
          ]}
        />
        <NavItem to={holdings}>
          Holdings
        </NavItem>
        <NavDropdown
          title='Funds'
          navItems={[
            {to: funds, title: 'Funds'},
            {to: assetCategories, title: 'Asset Categories'},
          ]}
        />
        <NavItem to={portfolios}>
          Portfolios
        </NavItem>
      </NavBar>
      <Switch>
        <Route path='/' exact component={HomePage}/>
        <Route path={accounts} component={AccountsPage}/>
        <Route path={accountCategories} component={AccountCategoriesPage}/>
        <Route path={contributions} component={ContributionsPage}/>
        <Route path={incomes} component={IncomesPage}/>
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
