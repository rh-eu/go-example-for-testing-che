import React, { Component } from "react";


class App extends Component {

  render() {

    return (

        <div>

          <h1>Data exposed by the Server:</h1>
          <p>------</p>
          <h2>{this.props.serverdata.hostname}</h2>
          <div>Demo application version <i>{this.props.serverdata.version}</i></div>

        </div>    
    );
  }
}

export default App;