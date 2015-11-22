import AppDispatcher from 'dispatchers/AppDispatcher';
import RouterContainer from 'services/RouterContainer';

import { LOGIN_USER, LOGOUT_USER } from 'constants';

var AuthActions  = {

  loginUser: (jwt) => {
    var savedJwt = localStorage.getItem('jwt');

    AppDispatcher.dispatch({
      actionType: LOGIN_USER,
      jwt: jwt
    });

    if (savedJwt !== jwt) {
      var nextPath = RouterContainer.get().getCurrentQuery().nextPath || '/';

      RouterContainer.get().transitionTo(nextPath);
      localStorage.setItem('jwt', jwt);
    }
  },

  logoutuser: () => {
    RouterContainer.get().transitionTo('/login');
    localStorage.removeItem('jwt');
    AppDispatcher.dispatch({
      actionType: LOGOUT_USER
    });
  }

};

export default AuthActions;
