import AuthActions from 'actions/AuthActions';
import request from 'reqwest';
import when from 'when';
import { LOGIN_URL, LOGOUT_URl } from 'constants';

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
}

export default new AuthService();
