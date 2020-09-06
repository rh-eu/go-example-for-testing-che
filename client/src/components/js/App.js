import React, { Component } from "react";
import Mem from "./Mem";

class App extends Component {

  constructor(props) {
    super(props);

    this.state = {
      token: ''
    };
  }

  getToken = () => {
    const data = { name: process.env.APIUser };
    let url = process.env.TOKEN_URL;
    console.log("apiPath: ", url);
    fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': 'Basic ' + btoa(process.env.BASICAUTH_CREDENTIALS),
        'Token': this.state.token,
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(response => {
      this.setState(response)
      console.log("Response: ", response)
      console.log("Received Token: ", this.state.token)
    });
  }
  
    componentDidMount() {
    console.log("Fetching Token ...")
    this.getToken()
  }

  render() {

    return (

        <div>

          <h1>Data exposed by the Server:</h1>
          <p>------</p>

          <p>{ this.props.hostname }</p>
          <p>.....</p>
          <p>{ this.props.addrs }</p>
          <p>.....</p>          
          <p>{ this.props.requestDump }</p>

        <p>Token : { this.state.token }</p>

        <Mem apiPath={"/memory/api"} token={this.state.token} getToken={this.getToken} />
           

        </div>

    );
  }
}

export default App;