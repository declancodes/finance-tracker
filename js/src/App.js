import React from "react";
import {hot} from "react-hot-loader"
import AccountPage from "./components/AccountPage";
import "./App.css";


class App extends React.Component {
  render() {
    return (
      <AccountPage />
    );
  }
}

export default hot(module)(App);