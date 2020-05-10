import React from "react";
import { hot } from "react-hot-loader";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import AccountsPage from "./components/accounts/AccountsPage";
import CategoryPage from "./components/categories/CategoryPage";
import ExpensesPage from "./components/expenses/ExpensesPage";
import "./App.css";

class App extends React.Component {
  render() {
    return (
      <Router>
        <div>
          <nav>
            <ul>
              <li>
                <Link to="/accounts">Accounts</Link>
              </li>
              <li>
                <Link to="/accountcategories">Account Categories</Link>
              </li>
              <li>
                <Link to="/expenses">Expenses</Link>
              </li>
              <li>
                <Link to="/expensecategories">Expense Categories</Link>
              </li>
            </ul>
          </nav>
          <Switch>
            <Route path="/accountcategories">
              <CategoryPage categoryType="Account"/>
            </Route>
            <Route path="/expensecategories">
              <CategoryPage categoryType="Expense"/>
            </Route>
            <Route path="/accounts">
              <AccountsPage/>
            </Route>
            <Route path="/expenses">
              <ExpensesPage/>
            </Route>
          </Switch>
        </div>
      </Router>
    );
  }
}

export default hot(module)(App);