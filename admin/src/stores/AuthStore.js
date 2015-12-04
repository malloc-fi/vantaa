import AuthActions from 'actions/AuthActions';
import AuthService from 'services/AuthService';
import BaseStore from 'stores/BaseStore';
import jwt_decode from 'jwt-decode';
import { LOGIN_USER, LOGOUT_USER } from 'constants';

class AuthStore extends BaseStore {

  constructor() {
    super();
    this.subscribe(() => this._registerToActions);
    this._user = null;
    this._jwt = '';
  }

  _registerToActions = (actions) => {
    switch(actions.actionType) {
      case LOGIN_USER:
        this._jwt = actions.jwt;
        this._user = jwt_decode(this._jwt);
        this.emitChange();
        break;

      case LOGOUT_USER:
        this._user = null;
        this.emitChange();
        break;

      default:
        break;
    }
  }

  get user() {
    return this._user;
  }

  get jwt() {
    return this._jwt;
  }

  isLoggedIn() {
    if (!!this._user) {
      AuthService.validateToken();
    }
    return !!this._user;
  }
}

export default new AuthStore();
