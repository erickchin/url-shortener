import React, { Component } from 'react';
import axios from 'axios';

export default class Home extends Component {
    state = {
        url: '',
        displaySuccess: false,
        shortenUrl: ''
    }

    handleChange = (e) => {
        this.setState({ url: e.target.value });
    }

    handleSubmit = (e) => {
        e.preventDefault();
    
        const request = {
            original_url: this.state.url
        };
    
        // Post request to submit a new url
        // Change domain
        axios.post('http://localhost:8080/submit', request)
          .then(res => {
            console.log(res);
            console.log(res.data);
            this.setState({ 
                displaySuccess: true,
                shortenUrl: res.data
             });
          })

      }

    render() {
        const displaySuccess = this.state.displaySuccess
        const shortenUrl = this.state.shortenUrl
        return (
            <div className="container">
                <h1 className="title">URL Shortener</h1>
                <p className="center-text">Enter the URL you want to shorten</p>
                <div className="center-text">
                    <form onSubmit={this.handleSubmit}>
                        <input placeholder="https:/google.ca" type="text" name="url" onChange={this.handleChange} />
                        <button type="submit">Submit</button>
                    </form>
                </div>
                {displaySuccess ? <p className="center-text">Success! Your link is: <a href={shortenUrl}>{shortenUrl}</a></p> : null}
            </div>
        );
    }
}