import React from 'react';
import {Route} from 'react-router';

import Main from 'components/Main/Main';
import SuperSecure from 'components/restricted';
import LoginForm from 'components/LoginForm/LoginForm';

const routes = (
  <Route handler={ Main }>
    <Route name="login" handler={ LoginForm } />
    <Route name="secure" handler={ SuperSecure } />
  </Route>
);

export default routes;
