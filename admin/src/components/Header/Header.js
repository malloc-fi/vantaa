import React from 'react';
import AuthStore from 'stores/AuthStore';
import { Route, Link } from 'react-router';
import AuthService from 'services/AuthService';

class Header extends React.Component {

  constructor(props) {
    super(props)
    this.state = this._getLoginState();
  }

  _getLoginState() {
    return {
      isLoggedIn: AuthStore.isLoggedIn()
    }
  }

  componentDidMount() {
    this.changeListener = this._onChange;
    AuthStore.addChangeListener(this.changeListener);
  }

  _onChange = () => {
    this.setState(this._getLoginState);
  }

  componentWillUnmount() {
    AuthStore.removeChangeListener(this.changeListener);
  }

  _logout(e) {
    e.preventDefault();
    AuthService.logout();
  }

  _contextualHeader = () => {
    if (this.state.isLoggedIn) {
      return(
        <ul className="nav header-nav">
          <li><a href="" onClick={ this._logout }>Logout</a></li>
        </ul>
      );
    } else {
      return(
        <ul className="nav header-nav">
          <li><Link to="login">Login</Link></li>
        </ul>
      );
    }
  }

  render() {
    return(
      <section id="header" className="header">
        { this._contextualHeader() }
      </section>
    );
  }
}

export default Header;
