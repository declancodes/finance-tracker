import React from "react";
import {hot} from "react-hot-loader"
import AccountsPage from "./components/accounts/AccountsPage";
import CategoryPage from "./components/common/CategoryPage";
import "./App.css";

class App extends React.Component {
  render() {
    return (
      <div>
        <CategoryPage categoryType="Account"/>
        <CategoryPage categoryType="Expense"/>
        <AccountsPage/>
      </div>
    );
  }
}

export default hot(module)(App);