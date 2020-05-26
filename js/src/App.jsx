import React from 'react';
import { hot } from 'react-hot-loader';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from 'react-router-dom';
import { AccountsPage } from './components/pages/AccountsPage';
import { CategoriesPage } from './components/pages/CategoriesPage';
import { ContributionsPage } from './components/pages/ContributionsPage';
import { ExpensesPage } from './components/pages/ExpensesPage';
import { FundsPage } from './components/pages/FundsPage';
import { HoldingsPage } from './components/pages/HoldingsPage';
import './App.css';

const App = () => (
  <Router>
    <div>
      <nav>
        <ul>
          <li><Link to='/accounts'>Accounts</Link></li>
          <li><Link to='/accountcategories'>Account Categories</Link></li>
          <li><Link to='/contributions'>Contributions</Link></li>
          <li><Link to='/expenses'>Expenses</Link></li>
          <li><Link to='/expensecategories'>Expense Categories</Link></li>
          <li><Link to='/funds'>Funds</Link></li>
          <li><Link to='/holdings'>Holdings</Link></li>
        </ul>
      </nav>
      <Switch>
        <Route path='/accounts'>
          <AccountsPage/>
        </Route>
        <Route path='/accountcategories'>
          <CategoriesPage key='1' categoryType='Account'/>
        </Route>
        <Route path='/contributions'>
          <ContributionsPage/>
        </Route>
        <Route path='/expenses'>
          <ExpensesPage/>
        </Route>
        <Route path='/expensecategories'>
          <CategoriesPage key='2' categoryType='Expense'/>
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