require('./loginForm.css');

import connectToStores from 'alt/utils/connectToStores';  
import AuthService from 'services/AuthService';
import React from 'react';
import { LinkedState } from 'decorators/react/LinkedState';

@LinkedState
class LoginForm extends React.Component {

  constructor() {
    super();
      this.state = {
        email: '',
        password: ''
    };
  }

  login = e => {
    e.preventDefault();
    AuthService.login(this.state.email, this.state.password)
      .catch(function(err) {
        console.log("Error logging in", err);
      });
  }

  render() {
    return(
      <form role="form">
        <div className="form-group">
          <input type="text" valueLink={ this.linkState('email') } placeholder="Username" />
          <input type="password" valueLink={ this.linkState('password') } placeholder="Password" />
        </div>
        <button type="submit" onClick={ this.login }>Submi</button>
      </form>
    );
  }
}

export default LoginForm;
