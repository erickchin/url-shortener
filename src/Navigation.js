import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

class Navigation extends Component {
  render() {
    return (
    <div className="black"><nav>
      <ul className="nav">
        <li className="nav">
         <Link to="/">Home</Link>
        </li>
        <li className="nav">
          <Link to="/logs">URL Logs</Link>
        </li>
      </ul>
    </nav></div>
      );
  }
}

export default Navigation;
