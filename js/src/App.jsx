import React from 'react';
import {
  Navbar,
  Nav,
  NavDropdown
} from 'react-bootstrap';
import { hot } from 'react-hot-loader';
import {
  BrowserRouter,
  Switch,
  Route,
} from 'react-router-dom';
import { LinkContainer } from 'react-router-bootstrap'
import { AccountsPage } from './components/pages/AccountsPage';
import { AccountCategoriesPage } from './components/pages/categories/AccountCategoriesPage';
import { AssetCategoriesPage } from './components/pages/categories/AssetCategoriesPage';
import { ContributionsPage } from './components/pages/ContributionsPage';
import { ExpensesPage } from './components/pages/ExpensesPage';
import { ExpenseCategoriesPage } from './components/pages/categories/ExpenseCategoriesPage';
import { FundsPage } from './components/pages/FundsPage';
import { HoldingsPage } from './components/pages/HoldingsPage';
import { HomePage } from './components/pages/HomePage';

const accounts = '/accounts';
const accountCategories = '/accountcategories';
const contributions = '/contributions';

const expenses = '/expenses';
const expenseCategories = '/expensecategories';

const funds = '/funds';
const assetCategories = '/assetcategories';
const holdings = '/holdings';

const App = () => {
  return (
    <BrowserRouter>
      <Navbar bg='dark' variant='dark'>
        <Navbar.Brand href='/'>Finance Tracker</Navbar.Brand>
        <Nav className='mr-auto'>
          <NavDropdown title='Accounts'>
            <LinkContainer to={accounts}>
              <NavDropdown.Item>Accounts</NavDropdown.Item>
            </LinkContainer>
            <LinkContainer to={accountCategories}>
              <NavDropdown.Item>Account Categories</NavDropdown.Item>
            </LinkContainer>
          </NavDropdown>
          <LinkContainer to={contributions}>
            <Nav.Link>Contributions</Nav.Link>
          </LinkContainer>
          <NavDropdown title='Expenses'>
            <LinkContainer to={expenses}>
              <NavDropdown.Item>Expenses</NavDropdown.Item>
            </LinkContainer>
            <LinkContainer to={expenseCategories}>
              <NavDropdown.Item>Expense Categories</NavDropdown.Item>
            </LinkContainer>
          </NavDropdown>
          <LinkContainer to={holdings}>
            <Nav.Link>Holdings</Nav.Link>
          </LinkContainer>
          <NavDropdown title='Funds'>
            <LinkContainer to={funds}>
              <NavDropdown.Item>Funds</NavDropdown.Item>
            </LinkContainer>
            <LinkContainer to={assetCategories}>
              <NavDropdown.Item>Asset Categories</NavDropdown.Item>
            </LinkContainer>
          </NavDropdown>
        </Nav>
      </Navbar>
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
      </Switch>
    </BrowserRouter>
  );
};

export default hot(module)(App);