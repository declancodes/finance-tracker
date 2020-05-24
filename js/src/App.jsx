import React from "react";
import { hot } from "react-hot-loader";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { AccountsPage } from "./components/pages/AccountsPage";
import { CategoryPage } from "./components/pages/CategoryPage";
import { ContributionsPage } from "./components/pages/ContributionsPage";
import { ExpensesPage } from "./components/pages/ExpensesPage";
import "./App.css";

const App = () => (
  <Router>
    <div>
      <nav>
        <ul>
          <li><Link to="/accounts">Accounts</Link></li>
          <li><Link to="/accountcategories">Account Categories</Link></li>
          <li><Link to="/contributions">Contributions</Link></li>
          <li><Link to="/expenses">Expenses</Link></li>
          <li><Link to="/expensecategories">Expense Categories</Link></li>
        </ul>
      </nav>
      <Switch>
        <Route path="/accounts">
          <AccountsPage/>
        </Route>
        <Route path="/accountcategories">
          <CategoryPage key="1" categoryType="Account"/>
        </Route>
        <Route path="/contributions">
          <ContributionsPage/>
        </Route>
        <Route path="/expenses">
          <ExpensesPage/>
        </Route>
        <Route path="/expensecategories">
          <CategoryPage key="2" categoryType="Expense"/>
        </Route>
      </Switch>
    </div>
  </Router>
);

export default hot(module)(App);