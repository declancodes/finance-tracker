import React from "react";
import {hot} from "react-hot-loader"
import CategoryPage from "./components/common/CategoryPage";
import "./App.css";

class App extends React.Component {
  render() {
    return (
      <div>
        <CategoryPage categoryType="Account"/>
        <CategoryPage categoryType="Expense"/>
      </div>
    );
  }
}

export default hot(module)(App);