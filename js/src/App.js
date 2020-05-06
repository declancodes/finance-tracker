import React, { Component } from "react";
import {hot} from "react-hot-loader"
import axios from "axios"
import Account from "./components/Account";
import "./App.css";


const API_URL = "http://localhost:8080"

class App extends Component {
  state = {
    accounts: []
  }

  componentDidMount() {
    const url = `${API_URL}/accounts`
    axios.get(url).then(response => response.data)
    .then((data) => {
      this.setState({ accounts: data })
      console.log(this.state.accounts)
    })
  }

  render() {
    return (
      <div className="accounts">
        {this.state.accounts.map(account =>
          (<Account account={account} key={account.accountUuid}/>)
        )}
      </div>
    );
  }
}

export default hot(module)(App);