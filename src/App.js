import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import Home from './home/Home'
import Logs from './logs/Logs'
import Navigation from './Navigation'

class App extends Component {
  render() {
    return (
      <Router>
        <Navigation/>
        <Route exact path='/' component={Home}/>
        <Route path='/logs' component={Logs}/>
      </Router>
    );
  }
}

export default App;
