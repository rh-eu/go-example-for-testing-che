import React, { Component } from "react";
import Mem from "./Mem";


class App extends Component {

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

          <Mem /> 

        </div>

    );
  }
}

export default App;