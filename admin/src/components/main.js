require('./main.css');

import React from 'react';
import { RouteHandler, Link } from 'react-router';
import LoginForm from 'components/loginForm/loginForm';

class Main extends React.Component {
  render() {
    return (
      <div>
        <h1>Home</h1>
        <LoginForm></LoginForm>
        <RouteHandler/>
      </div>
    );
  }
}

export default Main;
