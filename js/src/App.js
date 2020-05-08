import React from "react";
import {hot} from "react-hot-loader"
import AccountCategoryPage from "./components/AccountCategoryPage";
import "./App.css";


class App extends React.Component {
  render() {
    return (
      <AccountCategoryPage />
    );
  }
}

export default hot(module)(App);