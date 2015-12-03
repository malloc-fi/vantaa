import AuthActions from 'actions/AuthActions';
import AuthStore from 'stores/AuthStore';
import request from 'reqwest';
import when from 'when';
import {
  LOGIN_URL,
  LOGOUT_URL,
  VALIDATE_TOKEN_URL
} from 'constants';

class AuthService {

  login(email, password) {
    var data = {
      email: email,
      password: password
    };

    return this._handleAuth(when(request({
      url: LOGIN_URL,
      method: 'POST',
      crossOrigin: true,
      type: 'json',
      dataType: 'application/json',
      data: JSON.stringify(data)
    })));
  }

  _handleAuth(loginPromise) {
    return loginPromise
      .then(function(resp) {
        var jwt = resp.token;
        AuthActions.loginUser(jwt);
      });
  }

  logout() {
    var headersData = {};
    if (AuthStore.jwt) {
      headersData = {
        'Authorization': "Bearer " + AuthStore.jwt
      };
    }

    this._handleLogout(when(request({
      url: LOGOUT_URL,
      type: 'json',
      dataType: 'application/json',
      method: 'POST',
      crossOrigin: true,
      headers: {
        'Authorization': 'Bearer ' + AuthStore.jwt
      },
      error: function(e) {
        AuthActions.logoutUser();
      }
    })));
  }

  _handleLogout(logoutPromise) {
    logoutPromise.then(function(resp) {
      console.log(resp);
      AuthActions.logoutUser();
    });
  }

  validateToken() {
    var data = { token: AuthStore.jwt };
    return this._handleTokenValidation(when(request({
      url: VALIDATE_TOKEN_URL,
      method: 'POST',
      crossOrigin: true,
      type: 'json',
      dataType: 'application/json',
      data: data
    })));
  }

  _handleTokenValidation(validationPromise) {
    try {
      return validationPromise
      .then(function(resp) {
        return resp.status == 200;
      });
    }
    catch(e) {
      return false;
    }
  }
}

export default new AuthService();
