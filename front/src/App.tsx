import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';

const HEALTH_CHECK_ENDPOINT = ' http://localhost:1323/health_check';

interface IProps {}
interface IState {
  healthCheck: string,
}

export default class App extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);

    this.state = {
      healthCheck: ""
    }
  }

  handleHealthCheck() {
    axios
      .get(HEALTH_CHECK_ENDPOINT)
      .then((results) => {
        this.setState({healthCheck: results.data});
      })
      .catch(() => {
        this.setState({healthCheck: "failed health check"});
      });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.tsx</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          <br/>
          <p>{this.state.healthCheck}</p>
          <input
            type="button"
            value="Check Health"
            onClick={() => this.handleHealthCheck()}
          />
        </header>
      </div>
    );
  }
}