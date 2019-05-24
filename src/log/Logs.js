import React, { Component } from 'react';
import axios from 'axios';
import LogsTable from './LogsTable'

export default class Home extends Component {
    state = {
        url: '',
        displayInfo: false,
        logs: []
    }

    handleChange = (e) => {
        this.setState({ url: e.target.value });
    }

    handleSubmit = (e) => {
        e.preventDefault();
    
        var urlCode = this.state.url.split('/')[4]
        console.log(urlCode)
        axios.get("http://localhost:8080/log/" + urlCode)
        .then(res => {
            this.setState({ 
                displayInfo: true,
                logs: res.data
             });
        })
      }

    render() {
        const displayInfo = this.state.displayInfo
        const logs = this.state.logs
        return (
            <div className="container">
                <h1 className="title">URL Checker</h1>
                <p className="center-text">Enter the shortened URL you want to see</p>
                <div className="center-text">
                    <form onSubmit={this.handleSubmit}>
                        <input placeholder="https:/hostname.com/d/ASD2" type="text" name="url" onChange={this.handleChange} />
                        <button type="submit">Submit</button>
                    </form>
                </div>
            { displayInfo ? logs.length != 0 ? 
                    <div className="center-text">
                        <p >Accessed {logs.length} amount of times</p>
                        <LogsTable urlLogs={logs}/>
                    </div> 
                : <p className="center-text">There are no logs for this URL!</p>        
            : null}
                
            </div>
        );
    }
}