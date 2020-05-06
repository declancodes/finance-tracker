import React, { Component} from "react";
import {hot} from "react-hot-loader"
import axios from "axios"
import "./App.css";

const API_URL = "http://localhost:8080"

class App extends Component{
  state = {
    accounts: []
  }
  componentDidMount(){
    const url = `${API_URL}/accounts`
    axios.get(url).then(response => response.data)
    .then((data) => {
      this.setState({ accounts: data })
      console.log(this.state.accounts)
    })
  }
  render(){
    return(
      <div>
        <div>
          <h1>Finance Tracker</h1>
          <h2>Accounts</h2>
          {this.state.accounts.map((account) => (
            <div className="account" key={account.accountUuid}>
              <div className="account-body">
                <h5>{account.name}</h5>
                <h6>{account.accountCategory.name}</h6>
                <h6>{account.description}</h6>
                <h6>Amount: ${account.amount}</h6>
              </div>
            </div>
          ))}
        </div>
      </div>
   );
  }
}

export default hot(module)(App);