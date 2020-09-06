import React from 'react';
import {render} from 'react-dom';
import App from './components/js/App'

render(<App {...serverdata} />, document.getElementById("root"))
console.log("ServerData", serverdata)