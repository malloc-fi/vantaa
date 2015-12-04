var BASE_URL = 'http://localhost:9292';

export default {
  BASE_URL: BASE_URL,
  LOGIN_URL: BASE_URL + '/api/auth/token/new',
  LOGOUT_URL: BASE_URL + '/api/auth/logout',
  VALIDATE_TOKEN_URL: BASE_URL + '/api/auth/token/validate',

  // Actions
  LOGIN_USER: 'LOGIN_USER',
  LOGOUT_USER: 'LOGOUT_USER'
};
