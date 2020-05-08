import React from "react";
import {hot} from "react-hot-loader"
import AccountCategoryPage from "./components/accounts/AccountCategoryPage";
import ExpenseCategoryPage from "./components/expenses/ExpenseCategoryPage";
import "./App.css";

class App extends React.Component {
  render() {
    return (
      <div>
        <AccountCategoryPage />
        <ExpenseCategoryPage />
      </div>
    );
  }
}

export default hot(module)(App);