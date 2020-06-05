import React from 'react';
import { hot } from 'react-hot-loader';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom';
import { AccountsPage } from './components/pages/AccountsPage';
import { AccountCategoriesPage } from './components/pages/categories/AccountCategoriesPage';
import { AssetCategoriesPage } from './components/pages/categories/AssetCategoriesPage';
import { ContributionsPage } from './components/pages/ContributionsPage';
import { ExpensesPage } from './components/pages/ExpensesPage';
import { ExpenseCategoriesPage } from './components/pages/categories/ExpenseCategoriesPage';
import { FundsPage } from './components/pages/FundsPage';
import { HoldingsPage } from './components/pages/HoldingsPage';
import './App.css';

const App = () => (
  <Router>
    <div>
      <nav>
        <ul>
          <li>
            <ul>
              <li><Link to='/accounts'>Accounts</Link></li>
              <li><Link to='/accountcategories'>Account Categories</Link></li>
              <li><Link to='/contributions'>Contributions</Link></li>
            </ul>
          </li>
          <li>
            <ul>
              <li><Link to='/expenses'>Expenses</Link></li>
              <li><Link to='/expensecategories'>Expense Categories</Link></li>
            </ul>
          </li>
          <li>
            <ul>
              <li><Link to='/assetcategories'>Asset Categories</Link></li>
              <li><Link to='/funds'>Funds</Link></li>
              <li><Link to='/holdings'>Holdings</Link></li>
            </ul>
          </li>
        </ul>
      </nav>
      <Switch>
        <Route path='/accounts'>
          <AccountsPage/>
        </Route>
        <Route path='/accountcategories'>
          <AccountCategoriesPage/>
        </Route>
        <Route path='/assetcategories'>
          <AssetCategoriesPage/>
        </Route>
        <Route path='/contributions'>
          <ContributionsPage/>
        </Route>
        <Route path='/expenses'>
          <ExpensesPage/>
        </Route>
        <Route path='/expensecategories'>
          <ExpenseCategoriesPage/>
        </Route>
        <Route path='/funds'>
          <FundsPage/>
        </Route>
        <Route path='/holdings'>
          <HoldingsPage/>
        </Route>
      </Switch>
    </div>
  </Router>
);

export default hot(module)(App);